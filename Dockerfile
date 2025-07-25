FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY migrations/ ./migrations/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bot ./cmd

CMD ["/app/bot"]