include .env

migrateup:
	migrate -path ./internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@127.0.0.1:5432/${POSTGRESQL_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@127.0.0.1:5432/${POSTGRESQL_DB}?sslmode=disable" -verbose down

migrateup_server:
	migrate -path ./internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@127.0.0.1:10550/${POSTGRESQL_DB}?sslmode=disable" -verbose up

migratedown_server:
	migrate -path ./internal/repository/migrations -database "postgresql://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@127.0.0.1:10550/${POSTGRESQL_DB}?sslmode=disable" -verbose down


generateswag:
	swag init -g cmd/main.go

.PHONY: migrateup migratedown generateswag
