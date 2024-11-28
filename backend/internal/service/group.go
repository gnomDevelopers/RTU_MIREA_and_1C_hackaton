package service

import (
	"context"
	"database/sql"
	"errors"
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

	exists, err := s.repository.Exists(ctx, req.Name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if exists {
		return nil, errors.New("group already exists")
	}

	id, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &entities.CreateGroupResponse{
		ID: id,
	}

	return res, err
}

func (s *GroupService) GetAll(c context.Context, req *entities.GetGroupRequest) (*[]entities.Group, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	res, err := s.repository.GetAll(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return res, nil
}
