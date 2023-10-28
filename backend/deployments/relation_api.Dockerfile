FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY relation/api/etc/relation.yaml relation.yaml
COPY output/relation_api relation_api

ENTRYPOINT ["/relation_api", "-f", "relation.yaml"]
