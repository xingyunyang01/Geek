FROM golang:1.17-alpine AS build

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /build

COPY . .

RUN go build -o httpserver .

ENTRYPOINT  ["./httpserver"]
