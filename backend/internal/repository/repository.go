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

type UniversityRepository interface {
	Create(context.Context, string) (int, error)
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

type AuditoryRepository interface {
	Create(context.Context, []string) ([]int, error)
	GetById(context.Context, int) (*entities.Auditory, error)
	GetByType(context.Context, string) (*entities.Auditory, error)
	GetByCampusId(context.Context, int) (*entities.Auditory, error)
	GetByProfile(context.Context, string) (*entities.Auditory, error)
	GetByCapacity(context.Context, int) (*entities.Auditory, error)
	Update(context.Context, *entities.Auditory) error
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

type Repository struct {
	User       UserRepository
	University UniversityRepository
	Campus     CampusRepository
	Auditory   AuditoryRepository
	Class      ClassRepository
}
