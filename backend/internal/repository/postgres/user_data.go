package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type UserDataRepository struct {
	db *sql.DB
}

func NewUserDataRepository(db *sql.DB) *UserDataRepository {
	return &UserDataRepository{
		db: db,
	}
}

func (r *UserDataRepository) AddStudent(ctx context.Context, userData *entities.UserData) (int, error) {
	query := `
		INSERT INTO user_data (
			id, last_name, first_name, father_name, university_id, 
			permission_id, faculty_id, department_id, educational_direction
		) VALUES (
			$1, $2, $3, $4, $5, 
			$6, $7, $8, $9
		)
	`

	_, err := r.db.ExecContext(ctx, query, userData.ID, userData.LastName, userData.FirstName, userData.FatherName, userData.UniversityID, userData.PermissionID,
		userData.FacultyID, userData.DepartmentID, userData.EducationalDirection)
	if err != nil {
		return 0, err
	}

	return userData.ID, nil
}
