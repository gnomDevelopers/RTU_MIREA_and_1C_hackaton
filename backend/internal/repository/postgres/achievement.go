package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type AchievementRepository struct {
	db *sql.DB
}

func NewAchievementRepository(db *sql.DB) *AchievementRepository {
	return &AchievementRepository{db: db}
}

func (r *AchievementRepository) Create(ctx context.Context, achievement *entities.CreateAchievementRequest) (*entities.CreateAchievementResponse, error) {
	var id int
	query := `
		INSERT INTO achievement (name, result, user_id) VALUES ($1, $2, $3) RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, achievement.Name, achievement.Result, achievement.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &entities.CreateAchievementResponse{ID: id}, nil
}

func (r *AchievementRepository) GetById(ctx context.Context, id int) (*[]entities.Achievement, error) {
	query := `
		SELECT * FROM achievement WHERE id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	achievements := []entities.Achievement{}
	for rows.Next() {
		achievement := entities.Achievement{}
		err := rows.Scan(&achievement.ID, &achievement.Name, &achievement.Result, &achievement.UserID)
		if err != nil {
			return nil, err
		}
		achievements = append(achievements, achievement)
	}
	return &achievements, nil
}

func (r *AchievementRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM achievement WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
