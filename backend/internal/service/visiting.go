package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type VisitingService struct {
	repository repository.VisitingRepository
	timeout    time.Duration
}

func NewVisitingService(repository repository.VisitingRepository) *VisitingService {
	return &VisitingService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *VisitingService) Create(c context.Context, visiting *entities.Visiting) (int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	id, err := s.repository.Create(ctx, visiting)
	return id, err
}

func (s *VisitingService) GetByUserIdAndClassId(c context.Context, userID, classID int) (*entities.Visiting, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	visiting, err := s.repository.GetByUserIdAndClassId(ctx, userID, classID)
	return visiting, err
}

func (s *VisitingService) Update(c context.Context, visiting *entities.Visiting) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, visiting)
	if err != nil {
		return err
	}
	return nil
}