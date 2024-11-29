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
		timeout:    time.Duration(30) * time.Second,
	}
}

func (s *VisitingService) Add(c context.Context, visitings []entities.Visiting) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	v := visitings[0]

	ok, err := s.repository.ClassExist(ctx, v.ClassID)
	if err != nil {
		return err
	}
	if ok {
		for _, visiting := range visitings {
			err := s.repository.Update(ctx, &visiting)
			if err != nil {
				continue
			}
		}
	} else {
		for _, visiting := range visitings {
			_, err := s.repository.Create(ctx, &visiting)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func (s *VisitingService) Get(c context.Context, classID int) (*[]entities.VisitingInfo, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	ok, err := s.repository.ClassExist(ctx, classID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	res, err := s.repository.GetGroupVisiting(ctx, classID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *VisitingService) CheckIn(c context.Context, visiting *entities.CheckInRequest) error {
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
