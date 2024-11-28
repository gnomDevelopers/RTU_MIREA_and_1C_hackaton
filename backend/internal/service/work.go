package service

import (
	"context"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type WorkService struct {
	repository repository.WorkRepository
	timeout    time.Duration
	conf       *config.Config
}

func NewWorkService(repository repository.WorkRepository, conf *config.Config) *WorkService {
	return &WorkService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
		conf:       conf,
	}
}

func (s *WorkService) Create(c context.Context, workUser *entities.WorkUser) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Create(ctx, workUser)
	return err
}

func (s *WorkService) CreateHR(c context.Context, hr *entities.HR) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.CreateHR(ctx, hr)
	return err
}
