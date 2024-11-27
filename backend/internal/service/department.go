package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type DepartmentService struct {
	repository repository.DepartmentRepository
	timeout    time.Duration
}

func NewDepartmentService(repository repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}
func (s *DepartmentService) GetByUniversity(c context.Context, university string) (*[]entities.Department, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	departments, err := s.repository.GetByUniversity(ctx, university)
	return departments, err
}

func (s *DepartmentService) Create(c context.Context, req *entities.CreateDepartmentRequest) (*entities.CreateDepartmentResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	departmentID, err := s.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &entities.CreateDepartmentResponse{
		ID: departmentID,
	}

	return res, nil
}


func (s *UserService) GetByID(c context.Context, id int) (*entities.CreateDepartmentResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	depID, err := s.repository.(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entities.CreateDepartmentResponse{ID: depID}, nil
}