"""
Author: huangsihao7
Date: 2023-10-28 11:27:21
LastEditors: huangsihao7 1057434651@qq.com
LastEditTime: 2023-10-30 14:44:11
FilePath: /scooter-WSVA/ML/app/models/speech.py
Description: 
"""
from typing import List

from pydantic import BaseModel


class VideoSummaryKeyword(BaseModel):
    keywords: List[str]
    summary: str
    text: str
    duration: float
