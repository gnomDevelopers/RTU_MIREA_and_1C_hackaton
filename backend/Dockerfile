# сборка
FROM golang:1.22.5-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main ./cmd/main.go

# рантайм
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /main /app/main
ENTRYPOINT ["/app/main"]
