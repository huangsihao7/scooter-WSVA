FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY relation/rpc/etc/relation.yaml relation.yaml
COPY output/relation_rpc relation_rpc

ENTRYPOINT ["/relation_rpc", "-f", "relation.yaml"]
