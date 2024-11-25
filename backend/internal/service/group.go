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

func (s *GroupService) GetGroupMembers(c context.Context, groupName string) (*[]entities.GroupMember, error) {
	c, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	members, err := s.repository.GetGroupMembers(c, groupName)
	if err != nil {
		return nil, err
	}
	return members, nil
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

func (s *GroupService) Create(c context.Context, req *entities.CreateGroupRequest) (*entities.CreateGroupResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	//todo проверить no rows

	//exists, err := s.repository.Exists(ctx, req.Name)

	id, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &entities.CreateGroupResponse{
		ID: id,
	}

	return res, err
}
