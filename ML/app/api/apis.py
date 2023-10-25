from fastapi import APIRouter

from app.api.endpoints import speech

api_router = APIRouter()
api_router.include_router(speech.router, prefix="/speech", tags=["speech"])
