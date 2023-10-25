from contextlib import asynccontextmanager
import shutil
import logging
from fastapi import FastAPI

from app.core.logging import setup_logger
from app.api.apis import api_router
import uvicorn
import torch
from app.core.config import settings
from app.core.response import add_fastapi_middleware

logger = logging.getLogger("main")


@asynccontextmanager
async def lifespan(app: FastAPI):  # collects GPU memory
    logger.info(settings)
    yield
    if torch.cuda.is_available():
        torch.cuda.empty_cache()
        torch.cuda.ipc_collect()
    shutil.rmtree(settings.CACHE_ROOT_DIR)


description = """
骑上我心爱的小摩托，他永远不会堵车。
"""
app = FastAPI(title="骑上我心爱的小摩托", description=description, lifespan=lifespan)
app.include_router(api_router, prefix="/api/v1")
add_fastapi_middleware(app)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000, workers=1, log_config=setup_logger())
