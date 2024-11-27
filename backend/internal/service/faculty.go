package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type FacultyService struct {
	repositories repository.FacultyRepository
	timeout      time.Duration
}

func NewFacultyService(repo repository.FacultyRepository) *FacultyService {
	return &FacultyService{repo, time.Duration(10) * time.Second}
}

func (s *FacultyService) GetAll(c context.Context, req *entities.GetFacultyRequest) (*[]entities.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	faculties, err := s.repositories.GetAllByUniName(ctx, req)
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (s *FacultyService) Create(c context.Context, req *entities.CreateFacultyRequest) (*entities.CreateFacultyResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	res, err := s.repositories.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &entities.CreateFacultyResponse{ID: res}, err
}
