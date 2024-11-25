FROM golang:1.22.5-alpine AS builder

# Установка необходимых инструментов
RUN apk add --no-cache git make curl bash gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Установка утилит для генерации Swagger
RUN go install -ldflags="-s -w" github.com/swaggo/swag/cmd/swag@latest

# Сборка основного приложенияzw
RUN go build -ldflags="-s -w" -o /main ./cmd/main.go

RUN make generateswag

# Финальный образ
FROM alpine:3.18

# Установка зависимостей для работы приложения
RUN apk add --no-cache ca-certificates curl bash make


# Копирование приложения, окружения и makefile
COPY --from=builder /main /app/main
COPY --from=builder /app/.env /app/.env

WORKDIR /app

# Установка точки входа: выполнение makefile и запуск приложения
ENTRYPOINT ["sh", "-c", "/app/main || exit 1"]
