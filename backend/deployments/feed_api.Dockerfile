FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY feed/api/etc/feed-api.yaml feed-api.yaml
COPY output/feed_api feed_api

ENTRYPOINT ["/feed_api", "-f", "feed-api.yaml"]
