package postgres

import (
	"database/sql"
	"server/internal/repository"
)

func NewRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{
		User:         NewUserRepository(db),
		UserData:     NewUserDataRepository(db),
		University:   NewUniversityRepository(db),
		Campus:       NewCampusRepository(db),
		Class:        NewClassRepository(db),
		Audience:     NewAudienceRepository(db),
		UserSchedule: NewMyScheduleRepository(db),
	}
}
