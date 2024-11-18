package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type UniversityRepository struct {
	db *sql.DB
}

func NewUniversityRepository(db *sql.DB) *UniversityRepository {
	return &UniversityRepository{
		db: db,
	}
}

func (r *UniversityRepository) Create(ctx context.Context, name string) (int, error) {
	if name == "" {
		return 0, errors.New("")
	}

	query := `SELECT * FROM university WHERE name=$1`
	row := r.db.QueryRowContext(ctx, query, name)
	var tmp interface{}
	err := row.Scan(&tmp)
	if !(err == sql.ErrNoRows) {
		return 0, errors.New("class with these fields already exists")
	}

	var id int
	query = `INSERT INTO university (name) VALUES ($1) RETURNING id`

	err = r.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UniversityRepository) GetById(ctx context.Context, id int) (*entities.University, error) {
	var university entities.University
	query := `SELECT * FROM university WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&university.Id, &university.Name)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (r *UniversityRepository) GetByName(ctx context.Context, name string) (*entities.University, error) {
	var university entities.University
	query := `SELECT * FROM university WHERE name=$1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&university.Id, &university.Name)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (r *UniversityRepository) GetAll(ctx context.Context) (*[]entities.University, error) {
	var universities []entities.University
	query := `SELECT * FROM university`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var university entities.University
		err = rows.Scan(&university.Id, &university.Name)
		if err != nil {
			return nil, err
		}
		universities = append(universities, university)
	}

	return &universities, nil
}

func (r *UniversityRepository) Update(ctx context.Context, id int, name string) error {
	query := `UPDATE university SET name = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UniversityRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM university WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
