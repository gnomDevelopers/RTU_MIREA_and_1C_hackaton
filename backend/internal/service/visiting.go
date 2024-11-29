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

func (s *VisitingService) Update(c context.Context, visiting *entities.CheckInRequest) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	req := &entities.Visiting{
		UserID:  visiting.UserID,
		ClassID: visiting.ClassID,
		ID:      visiting.ID,
		Type:    "+",
	}
	err := s.repository.Update(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *VisitingService) CheckIn(c context.Context, visiting *entities.Visiting) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, visiting)
	if err != nil {
		return err
	}
	return nil
}

func (s *VisitingService) GetClassVisiting(c context.Context, classID int) (*[]entities.VisitingInfo, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	visits, err := s.repository.GetGroupVisiting(ctx, classID)
	if err != nil {
		return nil, err
	}
	return visits, nil
}
