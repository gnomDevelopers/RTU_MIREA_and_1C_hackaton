include .env

migrateup:
	migrate -path ./backend/internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@${POSTGRESQL_HOST}:${POSTGRESQL_PORT}/${POSTGRESQL_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path ./backend/internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@${POSTGRESQL_HOST}:${POSTGRESQL_PORT}/${POSTGRESQL_DB}?sslmode=disable" -verbose down

generateswag:
	swag init -g cmd/main.go

.PHONY: migrateup migratedown generateswag

