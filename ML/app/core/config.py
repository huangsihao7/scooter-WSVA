import tempfile

from pydantic import BaseSettings


class Settings(BaseSettings):
    WHISPER_MODEL_PATH: str = "openai/whisper-medium"
    WHISPER_PROMPT = """
- 我们的口号是什么？
- 骑上我心爱的小摩托，他永远不会堵车。
- 我们的目标是？
- 实习证明！
"""
    CACHE_ROOT_DIR = tempfile.mkdtemp()
    DEVICE = "cuda"

    class Config:
        case_sensitive = True


settings = Settings(_env_file=".env")  # type: ignore[call-arg]