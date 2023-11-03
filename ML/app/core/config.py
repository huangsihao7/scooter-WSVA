"""
Author: huangsihao7
Date: 2023-10-26 00:01:55
LastEditors: huangsihao7 1057434651@qq.com
LastEditTime: 2023-10-30 16:36:18
FilePath: /scooter-WSVA/ML/app/core/config.py
Description: 
"""
import tempfile
from typing import List

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    WHISPER_MODEL_PATH: str = "openai/whisper-medium"
    WHISPER_PROMPT: str = """
- 我们的口号是什么？\n- 骑上我心爱的小摩托，他永远不会堵车。\n- 我们的目标是？\n- 实习证明！
"""
    SUMMARY_PROMPT: str = """
您将获得一段视频内容的文本，您的任务是给出2个简体中文句子来总结视频。
下面是视频文本内容:
"""

    KEYWORD_TYPES: List[str] = [
        "美食",
        "动漫",
        "游戏",
        "网络小说",
        "科学",
        "编程",
        "文学",
        "化学",
        "计算机",
        "哲学",
        "经济学",
        "政治",
        "法律",
        "艺术",
        "健康",
        "音乐",
        "影视",
        "综艺",
        "直播",
        "明星",
        "搞笑",
        "脱口秀",
        "魔术",
        "舞蹈",
        "相声",
        "运动",
    ]

    KEYWORD_PROMPT: str = f"""
您将获得一段视频内容的文本，您的任务是从下面可选的标签列表中为视频选择5个简体中文标签，以吸引观众。
请直接输出这5个标签，下面是可选的标签的列表；
{"|".join(KEYWORD_TYPES)}
下面是视频文本内容:
"""

    SENTIMENT_AYALYSIS_PROMPT: str = """
您将获得一段文本，您的任务是判断文本的情感倾向，正向文本请返回正向，负向文本请返回负向。
请直接输出情感倾向；
下面是文本内容:
"""
    CACHE_ROOT_DIR: str = tempfile.mkdtemp()
    DEVICE: str = "cuda"

    SPARK_APPID: str
    SPARK_API_SECRET: str
    SPARK_API_KEY: str
    SPARK_URL: str = "ws://spark-api.xf-yun.com/v2.1/chat"

    class Config:
        case_sensitive = True


settings = Settings(_env_file=".env")  # type: ignore[call-arg]
