import tempfile

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    WHISPER_MODEL_PATH: str = "openai/whisper-medium"
    WHISPER_PROMPT: str = "- 我们的口号是什么？\n- 骑上我心爱的小摩托，他永远不会堵车。\n- 我们的目标是？\n- 实习证明！"
    CACHE_ROOT_DIR: str = tempfile.mkdtemp()
    DEVICE: str = "cuda"

    SPARK_APPID: str
    SPARK_API_SECRET: str
    SPARK_API_KEY: str
    SPARK_URL: str = "ws://spark-api.xf-yun.com/v2.1/chat"

    class Config:
        case_sensitive = True


settings = Settings(_env_file=".env")  # type: ignore[call-arg]
