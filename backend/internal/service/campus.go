package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type CampusService struct {
	repository repository.CampusRepository
	timeout    time.Duration
}

func NewCampusService(repository repository.CampusRepository) *CampusService {
	return &CampusService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *CampusService) Create(c context.Context, campus *[]entities.Campus) ([]int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	ids, err := s.repository.Create(ctx, campus)
	return ids, err
}

func (s *CampusService) GetById(c context.Context, id int) (*entities.Campus, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	campus, err := s.repository.GetById(ctx, id)
	return campus, err
}

func (s *CampusService) GetByAddress(c context.Context, address string) (*entities.Campus, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	campus, err := s.repository.GetByAddress(ctx, address)
	return campus, err
}

func (s *CampusService) GetByName(c context.Context, name string) (*entities.Campus, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	campus, err := s.repository.GetByName(ctx, name)
	return campus, err
}

func (s *CampusService) GetByUniversity(c context.Context, university string) (*[]entities.Campus, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	campuses, err := s.repository.GetByUniversity(ctx, university)
	return campuses, err
}

func (s *CampusService) GetAll(c context.Context) (*[]entities.Campus, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	campuses, err := s.repository.GetAll(ctx)
	return campuses, err
}

func (s *CampusService) Update(c context.Context, campus *entities.Campus) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, campus)
	return err
}

func (s *CampusService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	return err
}
