FROM golang:1.21-alpine AS build-env

LABEL maintainer="bondhan novandy<bondhan.novandy@gmail.com>"

ENV APP_NAME=hashmap-web
ENV GO111MODULE=on
ENV GOPRIVATE=github.com/bondhan
ENV TZ=Asia/Jakarta
ENV GIT_TERMINAL_PROMPT=0
ENV CGO_ENABLED=0

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    curl \
    busybox-extras \
    make \
    git \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk update && apk add --no-cache coreutils

WORKDIR /src

RUN mkdir -p /src/hashmap-web
COPY . /src/hashmap-web

WORKDIR /src/hashmap-web

RUN ls -lah

RUN go build -o ./bin/$APP_NAME main.go

# clean container
FROM alpine:latest

RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    tzdata

ENV TZ=Asia/Jakarta
ENV APP_NAME=hashmap-web

RUN mkdir /app
WORKDIR /app

COPY --from=build-env /src/hashmap-web/bin/$APP_NAME /app/$APP_NAME
COPY --from=build-env /src/hashmap-web/.env_sample /app/.env

# below for testing only
#COPY --from=build-env /src/hashmap-web/.env_example /app/.env

EXPOSE 3030

CMD /app/$APP_NAME