package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type GpaRepository struct {
	db *sql.DB
}

func NewGpaRepository(db *sql.DB) *GpaRepository {
	return &GpaRepository{
		db: db,
	}
}

func (r *GpaRepository) Exists(ctx context.Context, id int) (bool, error) {
	var exists int
	query := `SELECT 1 FROM gpa WHERE user_id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *GpaRepository) Create(ctx context.Context, id int) error {
	query := `INSERT INTO gpa (user_id, value) VALUES ($1, $2)`
	err := r.db.QueryRowContext(ctx, query, id, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *GpaRepository) Get(ctx context.Context, id int) (*entities.Gpa, error) {
	var gpa entities.Gpa
	query := `SELECT user_id, value FROM gpa WHERE user_id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&gpa.UserId, &gpa.Value)
	if err != nil {
		return nil, err
	}
	return &gpa, nil
}

func (r *GpaRepository) Update(ctx context.Context, id int, value float64) error {
	if check, _ := r.Exists(ctx, id); check != true {
		err := r.Create(ctx, id)
		if err != nil {
			return err
		}
	}
	query := `UPDATE gpa SET value = $1 WHERE user_id = $2`
	_, err := r.db.ExecContext(ctx, query, value, id)
	if err != nil {
		return err
	}
	return nil
}
