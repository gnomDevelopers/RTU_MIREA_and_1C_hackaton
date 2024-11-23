package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type UserScheduleService struct {
	repository repository.UserScheduleRepository
	timeout    time.Duration
}

func NewUserScheduleService(repository repository.UserScheduleRepository) *UserScheduleService {
	return &UserScheduleService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *UserScheduleService) Create(c context.Context, mySchedule *entities.UserSchedule) (int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	id, err := s.repository.Create(ctx, mySchedule)
	return id, err
}

func (s *UserScheduleService) GetByUserId(c context.Context, id int) (*[]entities.UserSchedule, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	mySchedule, err := s.repository.GetByUserId(ctx, id)
	return mySchedule, err
}

func (s *UserScheduleService) Update(c context.Context, mySchedule *entities.UserSchedule) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, mySchedule)
	return err
}

func (s *UserScheduleService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	return err
}
