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
