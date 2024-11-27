package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
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
	query := `SELECT 1 FROM hr WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, university).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *WorkRepository) Create(ctx context.Context, workUser *entities.WorkUser) error {
	if check, _ := r.Exists(ctx, workUser.Id); check != true {
		query := `INSERT INTO work_user (id, phone_number, telegram, skills, cv_path, hide_profile) VALUES ($1, $2, $3, $4, $5, $6)`
		_, err := r.db.ExecContext(ctx, query, workUser.Id, workUser.PhoneNumber, workUser.Telegram, workUser.Skills, workUser.CVPath, workUser.HideProfile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WorkRepository) CreateHR(ctx context.Context, hr *entities.HR) error {
	if check, _ := r.ExistsHR(ctx, hr.Email); check != true {
		query := `INSERT INTO hr (email, password, last_name, first_name, father_name, company, image_link) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := r.db.ExecContext(ctx, query, hr.Email, hr.Password, hr.LastName, hr.FirstName, hr.FatherName, hr.Company, hr.ImageLink)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WorkRepository) GetById(ctx context.Context, id int) (*entities.University, error) {
	var university entities.University

	query := `SELECT * FROM university WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&university.Id, &university.Name, &university.Postfix)
	if err != nil {
		return nil, err
	}
	return &university, nil
}
