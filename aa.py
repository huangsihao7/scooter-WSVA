'''
Author: Xu Ning
Date: 2023-10-30 23:25:38
LastEditors: Xu Ning
LastEditTime: 2023-10-30 23:36:20
Description: 
FilePath: \scooter-WSVA\aa.py
'''

# {
#     code:0,//成功
#     data:[
#         [20.188,0,16777215,'小明','好家伙，我直接好家伙'],
#         [1.188,0,16777215,'小华','好，我知道']
#     ]
#  }
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()


app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
data = {
    'code': 0,
    'data': [
            [20.188, 0, 16777215, '小明', '好家伙，我直接好家伙'],
            [0.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
            [2.188, 0, 16777215, '小华', '好，我知道'],
            [3.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
            [1.188, 0, 16777215, '小华', '好，我知道'],
    ]
}


@app.get('/aaa')
def aaa():
    return data


@app.get('/aaa/v3')
def aaa(id: str, max: int):
    return data
