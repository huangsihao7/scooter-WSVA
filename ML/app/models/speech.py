from typing import List

from pydantic import BaseModel


class VideoSummaryKeyword(BaseModel):
    keywords: List[str]
    summary: str
    text: str
