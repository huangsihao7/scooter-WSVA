FROM pytorch/pytorch:2.0.0-cuda11.7-cudnn8-runtime
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /workspace

RUN apt update && apt install -y ffmpeg
COPY requirements.txt /workspace/requirements.txt

RUN pip3 install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
COPY main.py /workspace/main.py
COPY .env /workspace/.env
COPY app /workspace/app

ENTRYPOINT ["python3","main.py"]
