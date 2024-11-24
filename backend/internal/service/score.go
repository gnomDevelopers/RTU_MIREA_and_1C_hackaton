package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type ScoreService struct {
	repository repository.ScoreRepository
	timeout    time.Duration
}

func NewScoreService(repository repository.ScoreRepository) *ScoreService {
	return &ScoreService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *ScoreService) Update(c context.Context, score *entities.Score) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, score)
	return err
}

func (s *ScoreService) Get(c context.Context, score *entities.Score) (*entities.Score, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	score, err := s.repository.Get(ctx, score)
	return score, err
}
