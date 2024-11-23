package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type GroupService struct {
	repository repository.GroupRepository
	timeout    time.Duration
}

func NewGroupService(repository repository.GroupRepository) *GroupService {
	return &GroupService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *GroupService) GetByUserID(c context.Context, userID int) (*entities.Group, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	group, err := s.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return group, nil
}
