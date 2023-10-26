import logging
from typing import List

import ffmpeg
import numpy as np
from fastapi import APIRouter
from pydantic import AnyHttpUrl
from torchvision.datasets.utils import download_url
from transformers import WhisperForConditionalGeneration, WhisperProcessor

from app.core.config import settings
from app.utils.sparkapi import sparkAPI
from app.models.speech import VideoSummaryKeyword

router = APIRouter()
logger = logging.getLogger("speech")

whisper_processor = WhisperProcessor.from_pretrained(settings.WHISPER_MODEL_PATH)
forced_decoder_ids = whisper_processor.get_decoder_prompt_ids(
    language="chinese", task="transcribe", no_timestamps=False
)
prompt_ids = whisper_processor.get_prompt_ids(
    settings.WHISPER_PROMPT, return_tensors="pt"
).to(settings.DEVICE)
whisper_model = WhisperForConditionalGeneration.from_pretrained(
    settings.WHISPER_MODEL_PATH
).to(settings.DEVICE)


def load_audio(file: str, sr: int = 16000) -> np.ndarray:
    try:
        out, _ = (
            ffmpeg.input(file, threads=0)
            .output("-", format="s16le", acodec="pcm_s16le", ac=1, ar=sr)
            .run(cmd=["ffmpeg", "-nostdin"], capture_stdout=True, capture_stderr=True)
        )
    except ffmpeg.Error as e:
        raise RuntimeError(f"Failed to load audio: {e.stderr.decode()}") from e

    return np.frombuffer(out, np.int16).flatten().astype(np.float32) / 32768.0


def _video2text(filename: str, url: str, sampling_rate=16000) -> str:
    """
    视频语音识别文字提取
    """
    download_url(url, settings.CACHE_ROOT_DIR, filename)
    audio = load_audio(f"{settings.CACHE_ROOT_DIR}/{filename}")
    input_features = whisper_processor(
        audio,
        sampling_rate=sampling_rate,
        return_tensors="pt",
    ).input_features
    # generate token ids
    predicted_ids = whisper_model.generate(
        input_features.to(settings.DEVICE),
        forced_decoder_ids=forced_decoder_ids,
        # prompt_ids=prompt_ids,
    )
    # decode token ids to text
    transcription = whisper_processor.batch_decode(
        predicted_ids, skip_special_tokens=True
    )
    split = transcription[0]
    try:
        split = transcription[0].split(settings.WHISPER_PROMPT)[-1]
    except Exception as e:
        pass
    logger.info(split)
    return split


def _textsummary(text: str) -> str:
    """
    文字摘要生成
    """
    summary = sparkAPI.ask(
        [
            {
                "role": "user",
                "content": settings.SUMMARY_PROMPT + text,
            }
        ]
    )
    logger.info(summary)
    return summary


def _textkeyword(text: str) -> List[str]:
    """
    文字摘要生成
    """
    keywords = sparkAPI.ask(
        [
            {
                "role": "user",
                "content": settings.KEYWORD_PROMPT + text,
            }
        ]
    )
    logger.info(keywords)
    try:
        keyword = keywords.split("\n")[-5:]
    except Exception as e:
        keyword = [keywords]
    return keyword


@router.get("/video2text", response_model=str)
async def video2text(
    *,
    video_url: AnyHttpUrl,
):
    """
    视频语音转文字
    """
    filename = video_url.__str__().split("/")[-1]
    url = video_url.__str__()
    return _video2text(filename, url)


@router.get("/video2summary", response_model=str)
async def video2summary(
    *,
    video_url: AnyHttpUrl,
):
    """
    视频摘要提取
    """
    filename = video_url.__str__().split("/")[-1]
    url = video_url.__str__()
    text = _video2text(filename, url)
    return _textsummary(text)


@router.get("/video2keyword", response_model=str)
async def video2keyword(
    *,
    video_url: AnyHttpUrl,
):
    """
    视频关键词提取
    """
    filename = video_url.__str__().split("/")[-1]
    url = video_url.__str__()
    text = _video2text(filename, url)
    return _textkeyword(text)


@router.get("/video2keywordAndSummary", response_model=VideoSummaryKeyword)
async def video2keywordAndSummary(
    *,
    video_url: AnyHttpUrl,
):
    """
    视频摘要和关键词提取
    """
    filename = video_url.__str__().split("/")[-1]
    url = video_url.__str__()
    text = _video2text(filename, url)
    keywords = _textkeyword(text)
    summary = _textsummary(text)
    return VideoSummaryKeyword(keywords=keywords, summary=summary, text=text)
