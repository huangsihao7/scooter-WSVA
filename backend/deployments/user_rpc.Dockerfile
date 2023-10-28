FROM alpine:3.17
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY user/rpc/etc/user.yaml user.yaml 
COPY output/user_rpc user_rpc

ENTRYPOINT ["/user_rpc", "-f", "user.yaml"]
