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
		Group:        NewGroupRepository(db),
		Campus:       NewCampusRepository(db),
		Class:        NewClassRepository(db),
		Audience:     NewAudienceRepository(db),
		UserSchedule: NewMyScheduleRepository(db),
		Department:   NewDepartmentRepository(db),
		Faculty:      NewFacultyRepository(db),
		Grade:        NewGradeRepository(db),
	}
}
