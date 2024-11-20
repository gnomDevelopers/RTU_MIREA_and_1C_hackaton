package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type DepartmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (r *DepartmentRepository) Exists(ctx context.Context, department *entities.Department) (bool, error) {
	var exists int
	query := "SELECT 1 FROM department WHERE name = $1"
	err := r.db.QueryRowContext(ctx, query, department.Name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *DepartmentRepository) Create(ctx context.Context, department *entities.CreateDepartmentRequest) (*entities.CreateDepartmentResponse, error) {
	if department.Name == "" {
		return nil, errors.New("")
	}
	var id int
	query := `INSERT INTO department (name) VALUES ($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, department.Name).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &entities.CreateDepartmentResponse{ID: id}, nil
}

func (r *DepartmentRepository) GetById(ctx context.Context, id int) (*entities.Department, error) {
	department := &entities.Department{}
	query := `SELECT id, name FROM department WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&department.ID, &department.Name)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (r *DepartmentRepository) GetByName(ctx context.Context, name string) (*entities.Department, error) {
	department := &entities.Department{}
	query := `SELECT id, name FROM department WHERE name = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&department.ID, &department.Name)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (r *DepartmentRepository) GetAll(ctx context.Context) (*[]entities.Department, error) {
	var departments []entities.Department
	query := `SELECT id, name FROM department`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var department entities.Department
		err = rows.Scan(&department.ID, &department.Name)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	return &departments, nil
}

func (r *DepartmentRepository) Update(ctx context.Context, department *entities.UpdateDepartmentRequest) error {
	query := `UPDATE department SET name = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, department.Name, department.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *DepartmentRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM department WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}