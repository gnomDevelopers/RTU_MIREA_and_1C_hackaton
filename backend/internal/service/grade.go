package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type GradeService struct {
	repository repository.GradeRepository
	timeout    time.Duration
}

func NewGradeService(repository repository.GradeRepository) *GradeService {
	return &GradeService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *GradeService) Create(c context.Context, grade *entities.Grade) (int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	id, err := s.repository.Create(ctx, grade)
	return id, err
}

func (s *GradeService) GetByUserIdAndClassId(c context.Context, userId, classId int) (*entities.Grade, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	campus, err := s.repository.GetByUserIdAndClassId(ctx, userId, classId)
	return campus, err
}
