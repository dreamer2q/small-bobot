FROM golang:1.17-alpine AS builder

RUN go env -w GO111MODULE=auto \
    && go env -w CGO_ENABLED=0 \
    && go env -w GOPROXY=https://goproxy.cn,direct 

WORKDIR /build

COPY ./ .

RUN set -ex \
    && cd /build \
    && go build -ldflags "-s -w -extldflags '-static'" -o smbot 

FROM alpine:3

ENV TIME_ZONE=Asia/Shanghai

RUN apk --update add --no-cache tzdata ffmpeg \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

COPY --from=builder /build/smbot /usr/bin/smbot

RUN chmod +x /usr/bin/smbot

WORKDIR /data

CMD [ "/usr/bin/smbot" ]
