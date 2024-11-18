include .env

migrateup:
	migrate -path ./internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@${POSTGRESQL_HOST}:${POSTGRESQL_PORT}/${POSTGRESQL_DB}?sslmode=disable" -verbose up


migratedown:
	migrate -path ./internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@${POSTGRESQL_HOST}:${POSTGRESQL_PORT}/${POSTGRESQL_DB}?sslmode=disable" -verbose down

generateswag:
	swag init -g cmd/main.go

.PHONY: migrateup migratedown generateswag

