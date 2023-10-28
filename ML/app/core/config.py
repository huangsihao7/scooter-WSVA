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
        "火锅",
        "披萨",
        "寿司",
        "汉堡",
        "麻辣烫",
        "烤肉",
        "粤菜",
        "韩式炸鸡",
        "意大利面",
        "北京烤鸭",
        "冰淇淋",
        "烧烤",
        "红烧肉",
        "蛋糕",
        "咖啡",
        "小龙虾",
        "粥",
        "烤鱼",
        "酸辣粉",
        "烧饼",
        "动漫",
        "漫画",
        "游戏",
        "Cosplay",
        "虚拟偶像",
        "网络小说",
        "模型",
        "动画片",
        "二次元音乐",
        "文化衫",
        "卡牌游戏",
        "漫展",
        "二次元周边",
        "角色扮演",
        "番剧",
        "轻小说",
        "电子竞技",
        "影迷",
        "动漫歌曲",
        "二次元美图",
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
        "幽默",
        "模仿",
        "音乐会",
        "脱颖而出",
        "表演",
        "演唱会",
        "玩游戏",
        "相声",
        "美声",
        "广告",
        "玩具",
        "足球",
        "篮球",
        "网球",
        "游泳",
        "田径",
        "乒乓球",
        "羽毛球",
        "高尔夫",
        "滑雪",
        "排球",
        "棒球",
        "拳击",
        "自行车",
        "赛车",
        "登山",
        "曲棍球",
        "跳水",
        "健身",
        "马拉松",
        "滑板",
    ]

    KEYWORD_PROMPT: str = f"""
您将获得一段视频内容的文本，您的任务是为视频提供5个简体中文标签，以吸引观众。
请直接输出这5个标签，标签只能从以下标签列表中选择：
下面是可选的标签的列表；
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
