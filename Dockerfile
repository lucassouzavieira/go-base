FROM golang:1.18-alpine AS builder
RUN apk update && apk add make git gcc musl-dev gdb

ARG GITHUB_TOKEN
ARG WORK_DIR
ARG APP_DIR
ARG APP_NAME
ARG APP

ENV GITHUB_TOKEN=$GITHUB_TOKEN
ENV WORK_DIR=$WORK_DIR
ENV APP_DIR=$APP_DIR
ENV APP_NAME=$APP_NAME
ENV APP=$APP
ENV GO111MODULE=on

RUN apk add --update gcc g++ openssh git make

# Building steps
WORKDIR /go/src/github.com/lucassvieira/$APP

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make build
COPY build/$APP /app/$APP

FROM alpine:latest AS compile

RUN apk add gdb
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache \
    ca-certificates

COPY --from=builder /app/$APP /app/$APP

CMD ["./app/go-project-layout"]
