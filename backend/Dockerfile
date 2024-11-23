FROM golang:1.22.5-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /main ./cmd/main.go

FROM alpine:3.18

RUN apk add --no-cache ca-certificates curl && \
    adduser -D -g '' appuser

USER appuser

COPY --from=builder /main /app/main
COPY --from=builder /app/.env /app/.env

WORKDIR /app

ENTRYPOINT ["sh", "-c", "/app/main || exit 1"]
# ENTRYPOINT ["sh", "-c", "/app/main || tail -f /dev/null"]
# ENTRYPOINT ["/app/main"]
