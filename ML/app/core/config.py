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
        "漫画",
        "游戏",
        "网络小说",
        "科学",
        "历史",
        "数学",
        "文学",
        "物理",
        "生物",
        "化学",
        "计算机",
        "天文学",
        "心理学",
        "哲学",
        "经济学",
        "地理",
        "政治",
        "社会学",
        "文化",
        "法律",
        "音乐",
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
    CACHE_ROOT_DIR: str = tempfile.mkdtemp()
    DEVICE: str = "cuda"

    SPARK_APPID: str
    SPARK_API_SECRET: str
    SPARK_API_KEY: str
    SPARK_URL: str = "ws://spark-api.xf-yun.com/v2.1/chat"

    class Config:
        case_sensitive = True


settings = Settings(_env_file=".env")  # type: ignore[call-arg]
