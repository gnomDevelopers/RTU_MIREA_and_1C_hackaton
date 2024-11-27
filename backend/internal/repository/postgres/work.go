package postgres

import (
	"context"
	"database/sql"
)

type WorkRepository struct {
	db *sql.DB
}

func NewWorkRepository(db *sql.DB) *WorkRepository {
	return &WorkRepository{
		db: db,
	}
}

func (r *WorkRepository) Exists(ctx context.Context, id int) (bool, error) {
	var exists int
	query := `SELECT 1 FROM work_user WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *WorkRepository) ExistsHR(ctx context.Context, university string) (bool, error) {
	var exists int
	query := `SELECT 1 FROM hr WHERE last_name = $1`
	err := r.db.QueryRowContext(ctx, query, university).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

//func (r *WorkRepository) Create(ctx context.Context, university *entities.CreateUniversityRequest) (int, error) {
//	if university.Name == "" {
//		return 0, errors.New("")
//	}
//	var id int
//	query := `INSERT INTO university (name, postfix) VALUES ($1, $2) RETURNING id`
//
//	err := r.db.QueryRowContext(ctx, query, university.Name, university.Postfix).Scan(&id)
//
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
//
//func (r *WorkRepository) CreateHR(ctx context.Context, university *entities.CreateUniversityRequest) (int, error) {
//	if university.Name == "" {
//		return 0, errors.New("")
//	}
//	var id int
//	query := `INSERT INTO university (name, postfix) VALUES ($1, $2) RETURNING id`
//
//	err := r.db.QueryRowContext(ctx, query, university.Name, university.Postfix).Scan(&id)
//
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
//
//func (r *WorkRepository) GetById(ctx context.Context, id int) (*entities.University, error) {
//	var university entities.University
//
//	query := `SELECT * FROM university WHERE id = $1`
//	err := r.db.QueryRowContext(ctx, query, id).Scan(&university.Id, &university.Name, &university.Postfix)
//	if err != nil {
//		return nil, err
//	}
//	return &university, nil
//}
