package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetById(ctx context.Context, id int64) (*entities.User, error) {
	user := entities.User{}
	query := `SELECT id, email, password FROM users WHERE id=$1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, login string) (*entities.User, error) {
	user := entities.User{}
	query := "SELECT id, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Password)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil

}

func (r *UserRepository) Exists(ctx context.Context, email string) (bool, error) {
	exists := 0
	query := "SELECT 1 FROM users WHERE email = $1 LIMIT 1"

	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var lastInsertId int
	query := `INSERT INTO users (email, password)
	VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &entities.User{}, err
	}
	user.ID = lastInsertId
	return user, nil
}
