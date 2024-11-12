# Этап, на котором выполняется сборка приложения
FROM golang:1.22.5-alpine as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /main ./cmd/main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]