package service

import (
	"context"
	"errors"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type UniversityService struct {
	repository repository.UniversityRepository
	timeout    time.Duration
}

func NewUniversityService(repository repository.UniversityRepository) *UniversityService {
	return &UniversityService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (u *UniversityService) Create(c context.Context, request *entities.CreateUniversityRequest) (int, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	exists, err := u.repository.Exists(ctx, request.Name)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("university already exists")
	}

	university := &entities.University{
		Name: request.Name,
	}

	id, err := u.repository.Create(ctx, university)
	return id, err
}

func (u *UniversityService) GetById(c context.Context, id int) (*entities.University, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	university, err := u.repository.GetById(ctx, id)
	return university, err
}

func (u *UniversityService) GetByName(c context.Context, name string) (*entities.University, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	university, err := u.repository.GetByName(ctx, name)
	return university, err
}

func (u *UniversityService) GetAll(c context.Context) (*[]entities.University, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	universities, err := u.repository.GetAll(ctx)
	return universities, err
}

func (u *UniversityService) Update(c context.Context, id int, name string) error {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	err := u.repository.Update(ctx, id, name)
	return err
}

func (u *UniversityService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	err := u.repository.Delete(ctx, id)
	return err
}
