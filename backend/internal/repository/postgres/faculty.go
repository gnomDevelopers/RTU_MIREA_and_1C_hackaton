package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type FacultyRepository struct {
	db *sql.DB
}

func NewFacultyRepository(db *sql.DB) *FacultyRepository {
	return &FacultyRepository{
		db: db,
	}
}

func (r *FacultyRepository) Create(ctx context.Context, faculty *entities.CreateFacultyRequest) (int, error) {
	if faculty.Name == "" {
		return 0, errors.New("")
	}
	query := `SELECT * FROM faculty WHERE name=$1`
	row := r.db.QueryRowContext(ctx, query, faculty.Name)
	var tmp interface{}
	err := row.Scan(&tmp)
	if !(err == sql.ErrNoRows) {
		return 0, errors.New("faculty with these fields already exists")
	}

	var id int
	query = `INSERT INTO faculty (name) VALUES ($1) RETURNING id`

	err = r.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *FacultyRepository) GetById(ctx context.Context, id int) (*entities.University, error) {
	var university entities.University
	query := `SELECT * FROM faculty WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&university.Id, &university.Name)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (r *FacultyRepository) GetByName(ctx context.Context, name string) (*entities.University, error) {
	var university entities.University
	query := `SELECT * FROM faculty WHERE name=$1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&university.Id, &university.Name)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (r *FacultyRepository) GetAll(ctx context.Context) (*[]entities.University, error) {
	var universities []entities.University
	query := `SELECT * FROM faculty`
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

func (r *FacultyRepository) Update(ctx context.Context, id int, name string) error {
	query := `UPDATE university SET name = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *FacultyRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM university WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
