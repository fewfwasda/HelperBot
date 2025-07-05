FROM golang:1.24.4-alpine
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o HelperBot.exe
CMD ["./HelperBot.exe"]