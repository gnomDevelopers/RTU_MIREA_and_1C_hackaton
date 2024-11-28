package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type AchievementService struct {
	repository repository.AchievementRepository
	timeout    time.Duration
}

func NewAchievementService(repository repository.AchievementRepository) *AchievementService {
	return &AchievementService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *AchievementService) GetByID(c context.Context, id int) (*[]entities.Achievement, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	res, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AchievementService) Create(c context.Context, req *entities.CreateAchievementRequest) (*entities.CreateAchievementResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	res, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AchievementService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
