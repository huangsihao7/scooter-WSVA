FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY feed/rpc/etc/feed.yaml feed.yaml
COPY output/feed_rpc feed_rpc

ENTRYPOINT ["/feed_rpc", "-f", "feed.yaml"]
