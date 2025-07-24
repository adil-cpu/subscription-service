FROM golang:1.23-alpine AS builder

WORKDIR /app

# устанавливаем git
RUN apk update && apk add --no-cache git

# копируем go.mod и go.sum, чтобы go mod tidy сработал корректно
COPY go.mod go.sum ./
RUN go mod tidy

# копируем все исходники, включая docs
COPY . .

# сборка
RUN go build -o main ./cmd/app

# финальный образ
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

EXPOSE 8080

CMD ["./main"]
