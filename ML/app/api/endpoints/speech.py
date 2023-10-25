import logging

import ffmpeg
import numpy as np
from fastapi import APIRouter
from pydantic import AnyHttpUrl
from torchvision.datasets.utils import download_url
from transformers import WhisperForConditionalGeneration, WhisperProcessor

from app.core.config import settings

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
        prompt_ids=prompt_ids,
    )
    # decode token ids to text
    transcription = whisper_processor.batch_decode(
        predicted_ids, skip_special_tokens=True
    )
    logger.info(transcription)
    return transcription[0]


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
