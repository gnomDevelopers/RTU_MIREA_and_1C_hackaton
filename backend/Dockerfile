FROM golang:1.22.5-alpine AS builder

# Установка необходимых инструментов
RUN apk add --no-cache git make curl bash gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Установка утилит для миграций и генерации Swagger
RUN go install -ldflags="-s -w" github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install -ldflags="-s -w" github.com/swaggo/swag/cmd/swag@latest && \
    go install -ldflags="-s -w" -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Сборка основного приложения
RUN go build -ldflags="-s -w" -o /main ./cmd/main.go

# Финальный образ
FROM alpine:3.18

# Установка зависимостей для работы приложения
RUN apk add --no-cache ca-certificates curl bash make


# Копирование приложения, окружения и makefile
COPY --from=builder /main /app/main
COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/makefile /app/makefile
COPY --from=builder /app/internal /app/internal

# Копирование утилит migrate и swag
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /go/bin/swag /usr/local/bin/swag

WORKDIR /app

# Установка точки входа: выполнение makefile и запуск приложения
ENTRYPOINT ["sh", "-c", "make generateswag migrateup && /app/main || exit 1"]
