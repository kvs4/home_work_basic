FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /http-sql

#-------------------------------------
FROM alpine:latest

COPY --from=builder /http-sql /http-sql

# Копируем файл .env из текущей директории на хосте в финальный контейнер
COPY --from=builder /app/.env .env

EXPOSE 8080

CMD ["/http-sql"]