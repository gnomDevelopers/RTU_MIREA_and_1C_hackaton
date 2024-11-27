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
	CreateAdmin(context.Context) error
	GetEducationalDirection(context.Context, int) (string, error)
	RefreshToken(int) (string, error)
	GetByID(context.Context, int) (*entities.User, error)
}

type University interface {
	Create(context.Context, *entities.CreateUniversityRequest) (int, error)
	GetById(context.Context, int) (*entities.University, error)
	GetByName(context.Context, string) (*entities.University, error)
	GetByUserId(context.Context, int) (*entities.University, error)
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

type Faculty interface {
	Create(c context.Context, req *entities.CreateFacultyRequest) (*entities.CreateFacultyResponse, error)
	GetAll(context.Context, *entities.GetFacultyRequest) (*[]entities.Faculty, error)
}

type Group interface {
	GetByUserID(context.Context, int) (*entities.Group, error)
	GetGroupMembers(context.Context, string) (*[]entities.GroupMember, error)
	Create(context.Context, *entities.CreateGroupRequest) (*entities.CreateGroupResponse, error)
}

type Service struct {
	UserService               *UserService
	UniversityService         *UniversityService
	CampusService             *CampusService
	GroupService              *GroupService
	AudienceService           *AudienceService
	ClassService              *ClassService
	UserScheduleService       *UserScheduleService
	GradeService              *GradeService
	ScoreService              *ScoreService
	FacultyService            *FacultyService
	AcademicDisciplineService *AcademicDisciplineService
	conf                      *config.Config
}

func NewService(repositories *repository.Repository, conf *config.Config) *Service {
	return &Service{
		UserService:               NewUserService(repositories.User, conf),
		UniversityService:         NewUniversityService(repositories.University),
		CampusService:             NewCampusService(repositories.Campus),
		ClassService:              NewClassService(repositories.Class),
		AudienceService:           NewAudienceService(repositories.Audience),
		GroupService:              NewGroupService(repositories.Group),
		UserScheduleService:       NewUserScheduleService(repositories.UserSchedule),
		GradeService:              NewGradeService(repositories.Grade),
		ScoreService:              NewScoreService(repositories.Score),
		FacultyService:            NewFacultyService(repositories.Faculty),
		AcademicDisciplineService: NewAcademicDisciplineService(repositories.AcademicDiscipline),
	}
}
