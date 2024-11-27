package service

import (
	"context"
	"errors"
	"fmt"
	"log"
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
	uniRepo    repository.UniversityRepository
	timeout    time.Duration
	conf       *config.Config
}

func NewUserService(repository repository.UserRepository, uniRepo repository.UniversityRepository, conf *config.Config) *UserService {
	return &UserService{
		repository: repository,
		uniRepo:    uniRepo,
		timeout:    time.Duration(10) * time.Second,
		conf:       conf,
	}
}

func (s *UserService) CreateUsers(ctx context.Context, requests []entities.CreateUserRequest) ([]entities.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var responses []entities.CreateUserResponse
	for _, request := range requests {
		group, err := s.uniRepo.GetById(ctx, request.UniversityID)
		if err != nil {
			return nil, err
		}
		postfix := group.Postfix

		fullName := util.GenerateLogin(fmt.Sprintf("%s %s %s", request.LastName, request.FirstName, request.FatherName))
		email := fmt.Sprintf("%s_%s@vuzplus.ru", fullName, postfix)

		exists, err := s.repository.Exists(ctx, email)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("user with email %s already exists", email)
		}

		hashedPassword, err := util.HashPassword(request.Password)
		if err != nil {
			return nil, err
		}

		u := &entities.User{
			Email:        email,
			Password:     hashedPassword,
			FirstName:    request.FirstName,
			LastName:     request.LastName,
			FatherName:   request.FatherName,
			Role:         request.Role,
			UniversityID: request.UniversityID,
		}

		switch request.Role {
		case "Учебный Отдел", "Декан":
			u.FacultyID = request.FacultyID
			u.DepartmentID = 1
			u.EducationalDirection = "null"
			u.GroupID = 1

		case "Заведующий кафедрой", "Преподаватель":
			u.FacultyID = request.FacultyID
			u.DepartmentID = request.DepartmentID
			u.EducationalDirection = "null"
			u.GroupID = 1

		case "Студент":
			u.FacultyID = request.FacultyID
			u.DepartmentID = request.DepartmentID
			u.EducationalDirection = request.EducationalDirection
			u.GroupID = request.GroupID
		}

		r, err := s.repository.CreateUser(ctx, u)
		if err != nil {
			return nil, err
		}

		responses = append(responses, entities.CreateUserResponse{
			ID: r.ID,
		})
	}

	return responses, nil
}

func (s *UserService) Login(c context.Context, request *entities.LoginUserRequest) (*entities.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.repository.GetByEmail(ctx, request.Email)
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

func (s *UserService) RefreshToken(userID int) (string, error) {
	TokenExpiration, err := strconv.Atoi(s.conf.Application.TokenExpiration)
	if err != nil {
		return "", errors.New("wrong data")
	}
	accessToken, err := pkg.GenerateAccessToken(userID, TokenExpiration, s.conf.Application.SigningKey)
	if err != nil {
		return "", errors.New("wrong data")
	}
	return accessToken, nil
}

func (s *UserService) CreateAdmin(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	exists, err := s.repository.Exists(ctx, "admin")
	if err != nil {
		return err
	}
	if exists {
		log.Println("admin already exists")
		return nil
	} else {
		log.Println("admin doesn't exist, creating")
	}

	password := util.GenerateTemporaryPassword(15)
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	createAdmin := &entities.User{
		Email:                "admin",
		Password:             hashedPassword,
		FirstName:            "Админ",
		LastName:             "Админов",
		FatherName:           "Админович",
		Role:                 "Администратор",
		UniversityID:         1,
		FacultyID:            1,
		GroupID:              1,
		DepartmentID:         1,
		EducationalDirection: "null",
	}

	user, err := s.repository.CreateUser(ctx, createAdmin)
	if err != nil {
		return err
	}

	log.Printf("Admin created\nemail: %s\npassword: %s", user.Email, password)
	return nil
}

func (s *UserService) GetEducationalDirection(c context.Context, userID int) (string, error) {
	user, err := s.repository.GetById(c, userID)
	if err != nil {
		return "", err
	}
	return user.EducationalDirection, nil
}

func (s *UserService) GetByUniversity(c context.Context, university string) (*[]entities.User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	users, err := s.repository.GetByUniversity(ctx, university)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetByID(c context.Context, userID int) (*entities.UserInfo, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.repository.GetInfoById(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
