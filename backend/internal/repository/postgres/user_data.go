package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type UserDataRepository struct {
	db *sql.DB
}

func NewUserDataRepository(db *sql.DB) *UserDataRepository {
	return &UserDataRepository{
		db: db,
	}
}

func (r *UserDataRepository) Add(context.Context, *entities.UserData) (int, error) {
	panic("implement me")
}
