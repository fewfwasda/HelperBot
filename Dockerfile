# Stage 1: сборка приложения
FROM golang:1.24.4-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARH=amd64 \ 
    go build -o helperbot

# Stage 2: минимальный рантайм
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/helperbot .

ENTRYPOINT [ "./helperbot" ]