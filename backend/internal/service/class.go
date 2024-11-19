package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type ClassService struct {
	repository repository.ClassRepository
	timeout    time.Duration
}

func NewClassService(repository repository.ClassRepository) *ClassService {
	return &ClassService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *ClassService) Create(c context.Context, class *[]entities.Class) ([]int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	ids, err := s.repository.Create(ctx, class)
	return ids, err
}

func (s *ClassService) GetById(c context.Context, id int) (*entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	class, err := s.repository.GetById(ctx, id)
	return class, err
}

func (s *ClassService) GetByGroupName(c context.Context, groupName string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByGroupName(ctx, groupName)
	return classes, err
}

func (s *ClassService) GetByTeacherName(c context.Context, teacherName string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByTeacherName(ctx, teacherName)
	return classes, err
}

func (s *ClassService) GetByAuditoryId(c context.Context, auditoryId int) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByAuditoryId(ctx, auditoryId)
	return classes, err
}

func (s *ClassService) Update(c context.Context, class *entities.Class) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, class)
	return err
}

func (s *ClassService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	return err
}
