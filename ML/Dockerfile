FROM pytorch/pytorch:2.0.0-cuda11.7-cudnn8-runtime
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /workspace

RUN apt update && apt install ffmpeg
COPY requirements.txt /workspace/requirements.txt

RUN pip3 install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple


ENTRYPOINT ["python3","main.py"]
