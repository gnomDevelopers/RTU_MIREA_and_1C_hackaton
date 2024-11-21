package service

import (
	"context"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/repository"
)

type User interface {
	CreateUser(context.Context, *entities.CreateUserRequest) (*entities.CreateUserResponse, error)
	Login(context.Context, *entities.LoginUserRequest) (*entities.LoginUserResponse, error)
}

type UserData interface {
	Add(context.Context, *entities.AddUserDataRequest) error
}

type University interface {
	Create(context.Context, *entities.CreateUniversityRequest) (int, error)
	GetById(context.Context, int) (*entities.University, error)
	GetByName(context.Context, string) (*entities.University, error)
	GetAll(context.Context) (*[]entities.University, error)
	Update(context.Context, int, string) error
	Delete(context.Context, int) error
}

type Campus interface {
	Create(context.Context, int, string, string) (int, error)
	GetById(context.Context, int) (*entities.Campus, error)
	GetByAddress(context.Context, string) (*entities.Campus, error)
	GetByName(context.Context, string) (*entities.Campus, error)
	GetByUniversityId(context.Context, int) (*entities.Campus, error)
	GetAll(context.Context) (*[]entities.Campus, error)
	Update(context.Context, int, string) error
	Delete(context.Context, int) error
}

// TODO Дополнить для других сервисов
type Service struct {
	UserService         *UserService
	UserData            *UserDataService
	UniversityService   *UniversityService
	CampusService       *CampusService
	AudienceService     *AudienceService
	ClassService        *ClassService
	UserScheduleService *UserScheduleService
	conf                *config.Config
}

func NewService(repositories *repository.Repository, conf *config.Config) *Service {
	return &Service{
		UserService:       NewUserService(repositories.User, conf),
		UserData:          NewUserDataService(repositories.User, repositories.UserData, repositories.Faculty, repositories.Department, repositories.University),
		UniversityService: NewUniversityService(repositories.University),
		CampusService:     NewCampusService(repositories.Campus),
		ClassService:      NewClassService(repositories.Class),
		AudienceService:   NewAudienceService(repositories.Audience),
	}
}
