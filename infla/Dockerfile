# とりあえずgolangのバージョンは最新を使用
FROM --platform=linux/x86_64 golang:1.19.3-alpine
# アップデートとgitのインストール
WORKDIR /gin-rest
RUN apk update && \
    apk add --update --no-cache git && \
    apk add --no-cache gcc && \
    apk add --no-cache musl-dev && \
    apk add --no-cache make && \
    wget https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz && \
    tar -zxvf migrate.linux-amd64.tar.gz && \
    mkdir -p /usr/local/bin && \
    mv migrate /usr/local/bin/migrate

