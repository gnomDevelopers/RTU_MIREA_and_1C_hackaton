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
	query := `SELECT FROM users (id, login, email, password) 
	WHERE id=$1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*entities.User, error) {
	user := entities.User{}
	query := "SELECT id, email, login, password, name, surname, third_name FROM users WHERE login = $1"
	err := r.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Email, &user.Login,
		&user.Password)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil

}

func (r *UserRepository) Exists(ctx context.Context, login string, email string) (bool, error) {
	exists := 0
	query := "SELECT 1 FROM users WHERE login=$1 OR email = $2 LIMIT 1"

	err := r.db.QueryRowContext(ctx, query, login, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var lastInsertId int
	query := `INSERT INTO users (login, email, password)
	VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Login, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &entities.User{}, err
	}
	user.ID = lastInsertId
	return user, nil
}
