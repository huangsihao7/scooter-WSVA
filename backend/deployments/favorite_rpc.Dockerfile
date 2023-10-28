FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY favorite/rpc/etc/favorite.yaml favorite.yaml
COPY output/favorite_rpc favorite_rpc

ENTRYPOINT ["/favorite_rpc", "-f", "favorite.yaml"]
