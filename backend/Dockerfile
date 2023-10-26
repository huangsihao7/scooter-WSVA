FROM golang:alpine as builder

WORKDIR /build

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn/,direct

RUN apk update --no-cache \
    && apk upgrade \
    && apk add --no-cache bash \
            bash-doc \
            bash-completion \
    && apk add --no-cache tzdata \
    && rm -rf /var/cache/apk/*

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy

COPY . .

ENTRYPOINT ["sh", "docker-entrypoint.sh"]