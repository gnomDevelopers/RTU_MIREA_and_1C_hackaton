package postgres

import (
	"context"
	"database/sql"
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

func (u *UniversityRepository) Create(ctx context.Context, name string) (int, error) {
	var id int
	query := `INSERT INTO university (name) VALUES ($1) RETURNING id`

	err := u.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UniversityRepository) GetById(ctx context.Context, id int) (*entities.University, error) {
	var university entities.University
	query := `SELECT * FROM university WHERE id = $1`
	err := u.db.QueryRowContext(ctx, query, id).Scan(&university.Id, &university.Name)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (u *UniversityRepository) GetByName(ctx context.Context, name string) (*entities.University, error) {
	var university entities.University
	query := `SELECT * FROM university WHERE name=$1`
	err := u.db.QueryRowContext(ctx, query, name).Scan(&university.Id, &university.Name)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (u *UniversityRepository) GetAll(ctx context.Context) (*[]entities.University, error) {
	var universities []entities.University
	query := `SELECT * FROM university`
	rows, err := u.db.QueryContext(ctx, query)
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

func (u *UniversityRepository) Update(ctx context.Context, id int, name string) error {
	query := `UPDATE university SET name = $1 WHERE id = $2`
	err := u.db.QueryRowContext(ctx, query, name, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (u *UniversityRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM university WHERE id = $1`
	err := u.db.QueryRowContext(ctx, query, id).Err()
	if err != nil {
		return err
	}
	return nil
}
