import _thread as thread
import base64
import hashlib
import hmac
import json
import ssl
from datetime import datetime
from time import mktime
from urllib.parse import urlencode, urlparse
from wsgiref.handlers import format_date_time

import websocket

from app.core.config import settings


class WsParam(object):
    # 初始化
    def __init__(self, APPID, APIKey, APISecret, Spark_url):
        self.APPID = APPID
        self.APIKey = APIKey
        self.APISecret = APISecret
        self.host = urlparse(Spark_url).netloc
        self.path = urlparse(Spark_url).path
        self.Spark_url = Spark_url

    # 生成url
    def create_url(self):
        # 生成RFC1123格式的时间戳
        now = datetime.now()
        date = format_date_time(mktime(now.timetuple()))

        # 拼接字符串
        signature_origin = "host: " + self.host + "\n"
        signature_origin += "date: " + date + "\n"
        signature_origin += "GET " + self.path + " HTTP/1.1"

        # 进行hmac-sha256进行加密
        signature_sha = hmac.new(
            self.APISecret.encode("utf-8"),
            signature_origin.encode("utf-8"),
            digestmod=hashlib.sha256,
        ).digest()

        signature_sha_base64 = base64.b64encode(signature_sha).decode(encoding="utf-8")

        authorization_origin = f'api_key="{self.APIKey}", algorithm="hmac-sha256", headers="host date request-line", signature="{signature_sha_base64}"'

        authorization = base64.b64encode(authorization_origin.encode("utf-8")).decode(
            encoding="utf-8"
        )

        # 将请求的鉴权参数组合为字典
        v = {"authorization": authorization, "date": date, "host": self.host}
        # 拼接鉴权参数，生成url
        url = self.Spark_url + "?" + urlencode(v)
        # 此处打印出建立连接时候的url,参考本demo的时候可取消上方打印的注释，比对相同参数时生成的url与自己代码生成的url是否一致
        return url


class SparkAPI(object):
    def __init__(self):
        self.answer = ""
        self.wsParam = WsParam(
            settings.SPARK_APPID,
            settings.SPARK_API_KEY,
            settings.SPARK_API_SECRET,
            settings.SPARK_URL,
        )

    def gen_params(self, appid, domain, question: list):
        """
        通过appid和用户的提问来生成请参数
        """
        data = {
            "header": {"app_id": appid, "uid": "1234"},
            "parameter": {
                "chat": {"domain": domain, "temperature": 0.5, "max_tokens": 2048}
            },
            "payload": {"message": {"text": question}},
        }
        return data

    def on_message(self, ws, message):
        # print(message)
        data = json.loads(message)
        code = data["header"]["code"]
        if code != 0:
            print(f"请求错误: {code}, {data}")
            ws.close()
        else:
            choices = data["payload"]["choices"]
            status = choices["status"]
            content = choices["text"][0]["content"]
            print(content, end="")
            self.answer += content
            # print(1)
            if status == 2:
                ws.close()

    # 收到websocket错误的处理
    def on_error(self, ws, error):
        print("### error:", error)

    # 收到websocket关闭的处理
    def on_close(self, ws, one, two):
        print(" ")

    # 收到websocket连接建立的处理
    def on_open(self, ws):
        thread.start_new_thread(self.run, (ws,))

    def run(self, ws, *args):
        data = json.dumps(
            self.gen_params(appid=ws.appid, domain=ws.domain, question=ws.question)
        )
        ws.send(data)

    def ask(self, question: list, domain: str = "generalv2") -> str:
        wsUrl = self.wsParam.create_url()
        websocket.enableTrace(False)
        ws = websocket.WebSocketApp(
            wsUrl,
            on_message=self.on_message,
            on_error=self.on_error,
            on_close=self.on_close,
            on_open=self.on_open,
        )
        ws.appid = settings.SPARK_APPID
        ws.question = question
        ws.domain = domain
        ws.run_forever(sslopt={"cert_reqs": ssl.CERT_NONE})
        answer = self.answer
        self.answer = ""
        return answer


sparkAPI = SparkAPI()
