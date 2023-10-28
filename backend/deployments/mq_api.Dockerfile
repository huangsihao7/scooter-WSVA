FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY mq/api/etc/mq-api.yaml mq-api.yaml
COPY output/mq_api mq_api

ENTRYPOINT ["/mq_api", "-f", "mq-api.yaml"]
