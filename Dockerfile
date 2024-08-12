FROM golang:1.22-alpine AS builder

ARG APP_PORT

WORKDIR /app

COPY . .

RUN go mod download \
    && GOARCH=amd64 go build -o /messager-web-app

FROM alpine:latest AS application

WORKDIR /root/

COPY --from=builder /messager-web-app .

EXPOSE ${APP_PORT}

ENTRYPOINT ["./messager-web-app"]