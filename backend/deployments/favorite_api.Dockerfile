FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY favorite/api/etc/favorite-api.yaml favorite-api.yaml
COPY output/favorite_api favorite_api

ENTRYPOINT ["/favorite_api", "-f", "favorite-api.yaml"]
