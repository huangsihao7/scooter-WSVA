FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY comment/rpc/etc/comment.yaml comment.yaml 
COPY output/comment_rpc comment_rpc

ENTRYPOINT ["/comment_rpc", "-f", "comment.yaml"]
