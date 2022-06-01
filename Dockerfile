FROM golang:1.18-alpine AS builder

ENV APP=go-base
ENV GO111MODULE=on

RUN apk update && apk add git make gdb
WORKDIR /go/src/github.com/lucassouzavieira/$APP/

# Handle dependencies
COPY go.mod go.sum ./
COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 go build -o /build/app -a ./cmd/app
RUN mv /build/app /app

# Final image
FROM alpine:latest AS compile

RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app

EXPOSE 8080
EXPOSE 8081

CMD ["./app"]