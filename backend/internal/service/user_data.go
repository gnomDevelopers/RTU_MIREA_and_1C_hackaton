package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/entities"
	"server/internal/repository"
	"server/util"
	"time"
)

type UserDataService struct {
	UserRepository     repository.UserRepository
	UserDataRepository repository.UserDataRepository
	GroupRepository    repository.GroupRepository
	timeout            time.Duration
}

func NewUserDataService(userRepo repository.UserRepository, userDataRepo repository.UserDataRepository, groupRepo repository.GroupRepository) *UserDataService {
	return &UserDataService{
		UserRepository:     userRepo,
		UserDataRepository: userDataRepo,
		GroupRepository:    groupRepo,
		timeout:            time.Duration(30) * time.Second,
	}
}

func (s *UserDataService) AddUser(c context.Context, requests *[]entities.AddUserDataRequest) (*[]entities.AddUserDataResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	var responses []entities.AddUserDataResponse
	for _, request := range *requests {
		fullName := fmt.Sprintf("%s %s %s", request.LastName, request.FirstName, request.FatherName)
		generatedEmail := util.GenerateLogin(fullName)
		password := util.GenerateTemporaryPassword(15)
		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			return nil, err
		}

		// Создаём пользователя
		newUser := &entities.User{
			Email:    generatedEmail,
			Password: hashedPassword,
		}

		ok, err := s.UserRepository.Exists(ctx, newUser.Email)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if ok {
			continue
		}

		user, err := s.UserRepository.CreateUser(ctx, newUser)
		if err != nil {
			return nil, err
		}

		userData := &entities.UserData{
			ID:           user.ID,
			LastName:     request.LastName,
			FirstName:    request.FirstName,
			FatherName:   request.FatherName,
			UniversityID: request.UniversityID,
			Role:         request.Role,
		}

		switch request.Role {
		case "Учебный отдел", "Декан":
			userData.FacultyID = request.FacultyID
		case "Заведующий кафедрой", "Преподаватель":
			userData.FacultyID = request.FacultyID
			userData.DepartmentID = request.DepartmentID
		case "Студент":
			userData.FacultyID = request.FacultyID
			userData.DepartmentID = request.DepartmentID
			userData.EducationalDirection = request.EducationalDirection
			userData.Group = request.Group
			joinGroup := &entities.CreateGroupRequest{
				Name:   request.Group,
				UserID: user.ID,
			}
			_, err = s.GroupRepository.Create(ctx, joinGroup)
			if err != nil {
				return nil, err
			}
		}

		_, err = s.UserDataRepository.AddUser(ctx, userData)

		responses = append(responses, entities.AddUserDataResponse{
			LastName:   request.LastName,
			FirstName:  request.FirstName,
			FatherName: request.FatherName,
			Email:      generatedEmail,
			Password:   password,
		})
	}

	return &responses, nil
}

func (s *UserDataService) AddAdmin(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	exists, err := s.UserRepository.Exists(ctx, "admin")
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	password := util.GenerateTemporaryPassword(15)
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}
	// Создаём пользователя
	newAdmin := &entities.User{
		Email:    "admin",
		Password: hashedPassword,
	}

	admin, err := s.UserRepository.CreateUser(ctx, newAdmin)
	if err != nil {
		return err
	}

	adminData := &entities.UserData{
		ID:         admin.ID,
		LastName:   "Админов",
		FirstName:  "Админ",
		FatherName: "Админович",
		Role:       "Администратор",
	}

	_, err = s.UserDataRepository.AddAdmin(ctx, adminData)
	if err != nil {
		return err
	}
	log.Println("admin added\nemail:%s password:%s", admin.Email, password)
	return nil
}
