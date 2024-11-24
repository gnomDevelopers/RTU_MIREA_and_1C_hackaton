package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type ScoreRepository struct {
	db *sql.DB
}

func NewScoreRepository(db *sql.DB) *ScoreRepository {
	return &ScoreRepository{
		db: db,
	}
}

func (r *ScoreRepository) Update(ctx context.Context, score *entities.Score) error {
	if check, _ := r.Exist(ctx, score); check != true {
		err := r.Create(ctx, score)
		if err != nil {
			return err
		}
	}

	var scoreSelected entities.Score
	query := `SELECT * FROM score WHERE user_id = $1 AND subject_name=$2`

	err := r.db.QueryRowContext(ctx, query, score.UserId, score.SubjectName).Scan(&scoreSelected.UserId, &scoreSelected.Sum, &scoreSelected.Count, &scoreSelected.SubjectName)
	if err != nil {
		return err
	}

	query = `UPDATE score SET sum=$1, count=$2 WHERE user_id = $3 AND subject_name=$4`

	_, err = r.db.ExecContext(ctx, query, score.Sum+scoreSelected.Sum, scoreSelected.Count+1, score.UserId, score.SubjectName)
	if err != nil {
		return err
	}

	return nil
}

func (r *ScoreRepository) Exist(ctx context.Context, score *entities.Score) (bool, error) {
	var exists int
	query := `SELECT 1 FROM score WHERE user_id = $1 AND subject_name=$2`
	err := r.db.QueryRowContext(ctx, query, score.UserId, score.SubjectName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *ScoreRepository) Create(ctx context.Context, score *entities.Score) error {
	var exists int
	query := `INSERT INTO score (user_id, sum, count, subject_name) VALUES($1, $2, $3, $4)`
	err := r.db.QueryRowContext(ctx, query, score.UserId, 0, 0, score.SubjectName).Scan(&exists)
	if err != nil {
		return err
	}
	return nil
}

func (r *ScoreRepository) Get(ctx context.Context, score *entities.Score) (*entities.Score, error) {
	var scoreSelected entities.Score
	query := `SELECT * FROM score WHERE user_id = $1 AND subject_name=$2`

	err := r.db.QueryRowContext(ctx, query, score.UserId, score.SubjectName).Scan(&scoreSelected.UserId, &scoreSelected.Sum, &scoreSelected.Count, &scoreSelected.SubjectName)
	if err != nil {
		return nil, err
	}

	return &scoreSelected, nil
}
