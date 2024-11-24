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

func (r *FacultyRepository) Exists(ctx context.Context, name string) (bool, error) {
	var exists int
	query := "SELECT 1 FROM faculty WHERE name = $1"
	err := r.db.QueryRowContext(ctx, query, name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *FacultyRepository) Create(ctx context.Context, faculty *entities.CreateFacultyRequest) (int, error) {
	if faculty.Name == "" {
		return 0, errors.New("")
	}
	var id int
	query := `INSERT INTO faculty (name, university) VALUES ($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, faculty.Name, faculty.University).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FacultyRepository) GetById(ctx context.Context, id int) (*entities.Faculty, error) {
	faculty := &entities.Faculty{}
	query := `SELECT id, name, university FROM faculty WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&faculty.ID, &faculty.Name, &faculty.University)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (r *FacultyRepository) GetByName(ctx context.Context, name string) (*entities.Faculty, error) {
	faculty := &entities.Faculty{}
	query := `SELECT id, name, university FROM faculty WHERE name = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&faculty.ID, &faculty.Name, &faculty.University)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (r *FacultyRepository) GetAll(ctx context.Context) (*[]entities.Faculty, error) {
	var faculties []entities.Faculty
	query := `SELECT id, name, university FROM faculty`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var faculty entities.Faculty
		err = rows.Scan(&faculty.ID, &faculty.Name, &faculty.University)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	return &faculties, nil
}

func (r *FacultyRepository) GetAllByUniName(ctx context.Context, request *entities.GetFacultyRequest) (*[]entities.Faculty, error) {
	var faculties []entities.Faculty
	query := `SELECT id, name, university FROM faculty WHERE university = $1`
	rows, err := r.db.QueryContext(ctx, query, request.UniversityName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var faculty entities.Faculty
		err = rows.Scan(&faculty.ID, &faculty.Name, &faculty.University)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	return &faculties, nil
}

func (r *FacultyRepository) Update(ctx context.Context, faculty *entities.UpdateFacultyRequest) error {
	query := `UPDATE faculty SET name = $1, university = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, faculty.Name, faculty.University, faculty.ID)
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
