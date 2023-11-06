FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY live/rpc/etc/live.yaml live.yaml
COPY output/live_rpc live_rpc

ENTRYPOINT ["/live_rpc", "-f", "live.yaml"]
