\package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type AcademicDisciplineService struct {
	repository repository.AcademicDisciplineRepository
	timeout    time.Duration
}

func NewAcademicDisciplineService(repository repository.AcademicDisciplineRepository) *AcademicDisciplineService {
	return &AcademicDisciplineService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *AcademicDisciplineService) Create(c context.Context, academicDiscipline *entities.AcademicDiscipline) (int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	id, err := s.repository.Create(ctx, academicDiscipline)
	return id, err
}

func (s *AcademicDisciplineService) GetByEducationalDirectionAndSemester(c context.Context, educationalDirection string, semester int) (*[]entities.AcademicDiscipline, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	academicDiscipline, err := s.repository.GetByEducationalDirectionAndSemester(ctx, educationalDirection, semester)
	return academicDiscipline, err
}

func (s *AcademicDisciplineService) GetByEducationalDirectionAndName(c context.Context, name, educationalDirection string) (*entities.AcademicDiscipline, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	academicDiscipline, err := s.repository.GetByEducationalDirectionAndName(ctx, name, educationalDirection)
	return academicDiscipline, err
}

func (s *AcademicDisciplineService) Update(c context.Context, academicDiscipline *entities.AcademicDiscipline) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, academicDiscipline)
	return err
}

func (s *AcademicDisciplineService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	return err
}
