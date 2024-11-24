package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type GradeRepository struct {
	db *sql.DB
}

func NewGradeRepository(db *sql.DB) *GradeRepository {
	return &GradeRepository{
		db: db,
	}
}

func (r *GradeRepository) Create(ctx context.Context, grade *entities.Grade) (int, error) {
	if check, _ := r.Exist(ctx, grade); check == true {
		return 0, errors.New("grade already exists")
	}

	var id int
	query := `INSERT INTO grade (class_id, user_id, value) VALUES($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, grade.ClassId, grade.UserId, grade.Value).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *GradeRepository) Exist(ctx context.Context, grade *entities.Grade) (bool, error) {
	var exists int
	query := `SELECT 1 FROM grade WHERE user_id = $1 AND class_id=$2`
	err := r.db.QueryRowContext(ctx, query, grade.UserId, grade.ClassId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *GradeRepository) GetByUserIdAndClassId(ctx context.Context, userId, classId int) (*entities.Grade, error) {
	var grade entities.Grade

	query := `SELECT * FROM grade WHERE user_id = $1 AND class_id = $2`
	err := r.db.QueryRowContext(ctx, query, userId, classId).Scan(&grade.Id, &grade.ClassId, &grade.UserId, &grade.Value)
	if err != nil {
		return nil, err
	}
	return &grade, nil
}
