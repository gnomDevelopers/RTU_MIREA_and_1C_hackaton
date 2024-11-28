package service

import (
	"context"
	"errors"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/repository"
	"server/pkg"
	"server/util"
	"time"
)

type WorkService struct {
	repository repository.WorkRepository
	conf       *config.Config
	timeout    time.Duration
}

func NewWorkService(repository repository.WorkRepository, conf *config.Config) *WorkService {
	return &WorkService{
		repository: repository,
		conf:       conf,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *WorkService) Create(c context.Context, workUser *entities.WorkUser) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Create(ctx, workUser)
	return err
}

func (s *WorkService) CreateHR(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hrs := []entities.HR{
		{
			Id:         1,
			Email:      "kovaleva_a_i_hr@vuzplus.ru",
			Password:   "kovalevA123",
			LastName:   "Ковалёва",
			FirstName:  "Анна",
			FatherName: "Игоревна",
			Company:    "1С",
			ImageLink:  "https://avatars.mds.yandex.net/i?id=40bc6365fc63a25dfe5da3d6825613c7e218fd18-12544206-images-thumbs&n=13",
		},
		{
			Id:         2,
			Email:      "smirnov_d_a_hr@vuzplus.ru",
			Password:   "smirnoV123",
			LastName:   "Смирнов",
			FirstName:  "Дмитрий",
			FatherName: "Алексеевич",
			Company:    "Группа Астра",
			ImageLink:  "https://habrastorage.org/getpro/habr/upload_files/77d/0ab/0b1/77d0ab0b1183979e7b4f1e72e1079c28.jpg",
		},
		{
			Id:         3,
			Email:      "petrova_o_s_hr@vuzplus.ru",
			Password:   "petrovA123",
			LastName:   "Петрова",
			FirstName:  "Ольга",
			FatherName: "Сергеевна",
			Company:    "Сбер",
			ImageLink:  "https://avatars.mds.yandex.net/i?id=95feab2c83fd8a7db54c6d3e454d391be0ce21c0-13101189-images-thumbs&n=13",
		},
	}

	for _, hr := range hrs {
		hashedPassword, err := util.HashPassword(hr.Password)
		hr.Password = hashedPassword
		err = s.repository.CreateHR(ctx, &hr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *WorkService) Login(c context.Context, request *entities.LoginUserRequest) (*entities.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.repository.GetByEmailHR(ctx, request.Email)
	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}

	err = util.CheckPassword(request.Password, u.Password)
	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}

	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}
	accessToken, err := pkg.GenerateAccessToken(u.Id, 1000,
		s.conf.Application.SigningKey)
	if err != nil {
		return &entities.LoginUserResponse{}, err
	}

	refreshToken, err := pkg.GenerateRefreshToken(u.Id, s.conf.Application.SigningKey)
	if err != nil {
		return &entities.LoginUserResponse{}, err
	}

	return &entities.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ID:           u.Id,
	}, nil
}

func (s *WorkService) GetByIdWorkUser(c context.Context, id int) (*entities.WorkUser, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	workUser, err := s.repository.GetByIdWorkUser(ctx, id)
	return workUser, err
}

func (s *WorkService) Exists(c context.Context, id int) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	flag, err := s.repository.Exists(ctx, id)
	return flag, err
}

func (s *WorkService) UpdateWorkUser(c context.Context, workUser *entities.WorkUserUpdateRequest) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.UpdateWorkUser(ctx, workUser)
	return err
}

func (s *WorkService) CreateResponse(c context.Context, response *entities.Response) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.CreateResponse(ctx, response)
	return err
}

func (s *WorkService) GetWorkUserResponses(c context.Context, workUserId int) (*[]entities.Response, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	responses, err := s.repository.GetWorkUserResponses(ctx, workUserId)
	return responses, err
}

func (s *WorkService) GetHRResponses(c context.Context, hrId int) (*[]entities.Response, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	responses, err := s.repository.GetHRResponses(ctx, hrId)
	return responses, err
}

func (s *WorkService) GetAllWorkUserId(c context.Context) (*[]entities.FullWorkUser, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	fullWorkUsers, err := s.repository.GetAllWorkUserId(ctx)
	return fullWorkUsers, err
}
