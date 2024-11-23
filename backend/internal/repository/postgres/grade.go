package postgres

import (
	"context"
	"database/sql"
	"fmt"
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
	var id int
	query := `INSERT INTO grade (class_id, user_id, value) VALUES($1, $2, $3) RETURNING id`
	fmt.Println(grade)
	err := r.db.QueryRowContext(ctx, query, grade.ClassId, grade.UserId, grade.Value).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
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