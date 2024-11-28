package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type VisitingRepository struct {
	db *sql.DB
}

func NewVisitingRepository(db *sql.DB) *VisitingRepository {
	return &VisitingRepository{db: db}
}

func (r *VisitingRepository) Exist(ctx context.Context, visiting *entities.Visiting) (bool, error) {
	var exists int
	query := `SELECT 1 FROM visiting WHERE user_id = $1 AND class_id=$2`
	err := r.db.QueryRowContext(ctx, query, visiting.UserID, visiting.ClassID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *VisitingRepository) Create(ctx context.Context, visiting *entities.Visiting) (int, error) {
	if check, _ := r.Exist(ctx, visiting); check == true {
		return 0, errors.New("grade already exists")
	}

	var id int
	query := `INSERT INTO visiting (class_id, user_id, type) VALUES($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, visiting.ClassID, visiting.UserID, visiting.Type).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *VisitingRepository) GetByUserIdAndClassId(ctx context.Context, userID, classID int) (*entities.Visiting, error) {
	var visiting entities.Visiting

	query := `SELECT * FROM grade WHERE user_id = $1 AND class_id = $2`
	err := r.db.QueryRowContext(ctx, query, userID, classID).Scan(&visiting.ID, &visiting.ClassID, &visiting.UserID, &visiting.Type)
	if err != nil {
		return nil, err
	}
	return &visiting, nil
}
