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

func (r *FacultyRepository) Exists(ctx context.Context, faculty *entities.Faculty) (bool, error) {
	var exists int
	query := "SELECT 1 FROM faculty WHERE name = $1"
	err := r.db.QueryRowContext(ctx, query, faculty.Name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *FacultyRepository) Create(ctx context.Context, faculty *entities.CreateFacultyRequest) (*entities.CreateFacultyResponse, error) {
	if faculty.Name == "" {
		return nil, errors.New("")
	}
	var id int
	query := `INSERT INTO faculty (name) VALUES ($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, faculty.Name).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &entities.CreateFacultyResponse{ID: id}, nil
}

func (r *FacultyRepository) GetById(ctx context.Context, id int) (*entities.Faculty, error) {
	faculty := &entities.Faculty{}
	query := `SELECT id, name FROM faculty WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&faculty.ID, &faculty.Name)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (r *FacultyRepository) GetByName(ctx context.Context, name string) (*entities.Faculty, error) {
	faculty := &entities.Faculty{}
	query := `SELECT id, name FROM faculty WHERE name = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&faculty.ID, &faculty.Name)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (r *FacultyRepository) GetAll(ctx context.Context) (*[]entities.Faculty, error) {
	var faculties []entities.Faculty
	query := `SELECT id, name FROM faculty`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var faculty entities.Faculty
		err = rows.Scan(&faculty.ID, &faculty.Name)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	return &faculties, nil
}

func (r *FacultyRepository) Update(ctx context.Context, faculty *entities.UpdateFacultyRequest) error {
	query := `UPDATE faculty SET name = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, faculty.Name, faculty.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *FacultyRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM faculty WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
