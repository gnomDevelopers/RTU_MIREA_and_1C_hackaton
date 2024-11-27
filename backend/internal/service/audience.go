package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type AudienceService struct {
	repository repository.AudienceRepository
	timeout    time.Duration
}

func NewAudienceService(repository repository.AudienceRepository) *AudienceService {
	return &AudienceService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *AudienceService) Create(c context.Context, audiences *[]entities.Audience) ([]int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	ids, err := s.repository.Create(ctx, audiences)
	return ids, err
}

func (s *AudienceService) GetById(c context.Context, id int) (*entities.Audience, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	audience, err := s.repository.GetById(ctx, id)
	return audience, err
}

func (s *AudienceService) GetByCampus(c context.Context, campus string) (*[]entities.Audience, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	audiences, err := s.repository.GetByCampus(ctx, campus)
	return audiences, err
}

func (s *AudienceService) GetByType(c context.Context, typeStr string) (*[]entities.Audience, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	audiences, err := s.repository.GetByType(ctx, typeStr)
	return audiences, err
}

func (s *AudienceService) GetByProfile(c context.Context, profile string) (*[]entities.Audience, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	audiences, err := s.repository.GetByProfile(ctx, profile)
	return audiences, err
}

func (s *AudienceService) GetByCapacity(c context.Context, capacity int) (*[]entities.Audience, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	audiences, err := s.repository.GetByCapacity(ctx, capacity)
	return audiences, err
}

func (s *AudienceService) GetByUniversity(c context.Context, university string) (*[]entities.Audience, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	audiences, err := s.repository.GetByUniversity(ctx, university)
	return audiences, err
}

func (s *AudienceService) Update(c context.Context, campus *entities.Audience) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, campus)
	return err
}

func (s *AudienceService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	return err
}
