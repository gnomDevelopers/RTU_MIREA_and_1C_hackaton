package service

import (
	"context"
	"errors"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/repository"
	"server/pkg"
	"server/util"
	"strconv"
	"time"
)

type UserService struct {
	repository repository.UserRepository
	timeout    time.Duration
	conf       *config.Config
}

func (s *UserService) CreateUser(c context.Context, request *entities.CreateUserRequest) (*entities.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	exists, err := s.repository.Exists(ctx, request.Login, request.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	u := &entities.User{
		Login:    request.Login,
		Email:    request.Email,
		Password: hashedPassword,
	}

	r, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	TokenExpiration, err := strconv.Atoi(s.conf.Application.TokenExpiration)
	if err != nil {
		return &entities.CreateUserResponse{}, errors.New("wrong data")
	}
	accessToken, err := pkg.GenerateAccessToken(u.ID, TokenExpiration,
		s.conf.Application.SigningKey)
	if err != nil {
		return &entities.CreateUserResponse{}, err
	}

	refreshToken, err := pkg.GenerateRefreshToken(u.ID, s.conf.Application.SigningKey)
	if err != nil {
		return &entities.CreateUserResponse{}, err
	}

	res := &entities.CreateUserResponse{
		ID:           r.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil

}

func (s *UserService) Login(c context.Context, request *entities.LoginUserRequest) (*entities.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.repository.GetByLogin(ctx, request.Login)
	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}
	err = util.CheckPassword(request.Password, u.Password)
	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}

	TokenExpiration, err := strconv.Atoi(s.conf.Application.TokenExpiration)
	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}
	accessToken, err := pkg.GenerateAccessToken(u.ID, TokenExpiration,
		s.conf.Application.SigningKey)
	if err != nil {
		return &entities.LoginUserResponse{}, err
	}

	refreshToken, err := pkg.GenerateRefreshToken(u.ID, s.conf.Application.SigningKey)
	if err != nil {
		return &entities.LoginUserResponse{}, err
	}

	return &entities.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ID:           u.ID,
	}, nil

}

func NewUserService(repository repository.UserRepository, conf *config.Config) *UserService {
	return &UserService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
		conf:       conf,
	}
}