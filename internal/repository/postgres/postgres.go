package postgres

import (
	"database/sql"
	"server/internal/repository"
)

func NewRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{
		User: NewUserRepository(db),
	}
	//TODO дополнить
}
