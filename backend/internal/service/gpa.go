package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type GpaService struct {
	repository repository.GpaRepository
	timeout    time.Duration
}

func NewGpaService(repository repository.GpaRepository) *GpaService {
	return &GpaService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *GpaService) Update(c context.Context, id int, value float64) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, id, value)
	return err
}

func (s *GpaService) GetById(c context.Context, id int) (*entities.Gpa, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	gpa, err := s.repository.Get(ctx, id)
	return gpa, err
}
