package repository

import (
	"context"
	"server/internal/entities"
)

type UserRepository interface {
	GetById(context.Context, int64) (*entities.User, error)
	GetByLogin(context.Context, string) (*entities.User, error)
	Exists(context.Context, string, string) (bool, error)
	CreateUser(context.Context, *entities.User) (*entities.User, error)
}

//type UserDataRepository interface {
//	Add(context.Context, *entities.UserData) error
//}

type UniversityRepository interface {
	Exists(context.Context, string) (bool, error)
	Create(context.Context, *entities.University) (int, error)
	GetById(context.Context, int) (*entities.University, error)
	GetByName(context.Context, string) (*entities.University, error)
	GetAll(context.Context) (*[]entities.University, error)
	Update(context.Context, int, string) error
	Delete(context.Context, int) error
}

type CampusRepository interface {
	Create(context.Context, *[]entities.Campus) ([]int, error)
	GetById(context.Context, int) (*entities.Campus, error)
	GetByAddress(context.Context, string) (*entities.Campus, error)
	GetByName(context.Context, string) (*entities.Campus, error)
	GetByUniversityId(context.Context, int) (*entities.Campus, error)
	GetAll(context.Context) (*[]entities.Campus, error)
	Update(context.Context, *entities.Campus) error
	Delete(context.Context, int) error
}

type AudienceRepository interface {
	Create(context.Context, *[]entities.Audience) ([]int, error)
	GetById(context.Context, int) (*entities.Audience, error)
	GetByType(context.Context, string) (*[]entities.Audience, error)
	GetByCampusId(context.Context, int) (*[]entities.Audience, error)
	GetByProfile(context.Context, string) (*[]entities.Audience, error)
	GetByCapacity(context.Context, int) (*[]entities.Audience, error)
	Update(context.Context, *entities.Audience) error
	Delete(context.Context, int) error
}

type ClassRepository interface {
	Create(context.Context, *[]entities.Class) ([]int, error)
	GetById(context.Context, int) (*entities.Class, error)
	GetByGroupName(context.Context, string) (*[]entities.Class, error)
	GetByTeacherName(context.Context, string) (*[]entities.Class, error)
	GetByAuditoryId(context.Context, int) (*[]entities.Class, error)
	Update(context.Context, *entities.Class) error
	Delete(context.Context, int) error
}

type FacultyRepository interface {
	Exists(context.Context, *entities.Faculty) (bool, error)
	Create(context.Context, *entities.CreateFacultyRequest) (*entities.CreateFacultyResponse, error)
	GetById(context.Context, int) (*entities.Faculty, error)
	GetByName(context.Context, string) (*entities.Faculty, error)
	GetAll(context.Context) (*[]entities.Faculty, error)
	Update(context.Context, *entities.UpdateFacultyRequest) error
	Delete(context.Context, int) error
}

type DepartmentRepository interface {
	Exists(context.Context, *entities.Department) (bool, error)
	Create(context.Context, *entities.CreateDepartmentRequest) (*entities.CreateDepartmentResponse, error)
	GetById(context.Context, int) (*entities.Department, error)
	GetByName(context.Context, string) (*entities.Department, error)
	GetAll(context.Context) (*[]entities.Department, error)
	Update(context.Context, *entities.UpdateDepartmentRequest) error
	Delete(context.Context, int) error
}

type Repository struct {
	User       UserRepository
	University UniversityRepository
	Campus     CampusRepository
	Audience   AudienceRepository
	Class      ClassRepository
	Faculty    FacultyRepository
	Department DepartmentRepository
}
