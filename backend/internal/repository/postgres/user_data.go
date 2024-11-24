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

func (r *UserDataRepository) AddUser(ctx context.Context, userData *entities.UserData) (int, error) {
	query := `
		INSERT INTO user_data (
			id, last_name, first_name, father_name, university_id, 
			role, faculty_id, department_id, educational_direction
		) VALUES (
			$1, $2, $3, $4, $5, 
			$6, $7, $8, $9
		)
	`

	_, err := r.db.ExecContext(ctx, query, userData.ID, userData.LastName, userData.FirstName, userData.FatherName, userData.UniversityID, userData.Role,
		userData.FacultyID, userData.DepartmentID, userData.EducationalDirection)
	if err != nil {
		return 0, err
	}

	return userData.ID, nil
}

func (r *UserDataRepository) AddAdmin(ctx context.Context, adminData *entities.UserData) (int, error) {
	query := `
		INSERT INTO user_data (
			id, last_name, first_name, father_name, role
		) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.ExecContext(ctx, query, adminData.ID, adminData.LastName, adminData.FirstName, adminData.FatherName, adminData.Role)
	if err != nil {
		return 0, err
	}

	return adminData.ID, nil
}

//GetAll(context.Context, int) (*[]entities.UserData, error)
//	GetAllByRole(context.Context, int, string) (*[]entities.UserData, error)

func (r *UserDataRepository) GetAll(ctx context.Context, universityID int) (*[]entities.GetUserDataResponse, error) {
	query := `
		SELECT ud.id, ud.last_name, ud.first_name, ud.father_name, ud.faculty_id, ud.department_id,
		ud.role, ud.educational_direction,g.name FROM user_data ud
		LEFT JOIN "group" g ON ud.id = g.user_id
		WHERE ud.university_id = :university_id 
    	AND ud.is_deleted = false;
	`

	rows, err := r.db.QueryContext(ctx, query, universityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userData []entities.GetUserDataResponse
	for rows.Next() {
		var user entities.GetUserDataResponse
		err = rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.FatherName, &user.FacultyID, &user.DepartmentID, &user.Role, &user.EducationalDirection, &user.Group)
		if err != nil {
			return nil, err
		}
		userData = append(userData, user)
	}
	return &userData, nil
}

func (r *UserDataRepository) GetAllByRole(ctx context.Context, universityID int, role string) (*[]entities.GetUserDataResponse, error) {
	query := `
		SELECT ud.id, ud.last_name, ud.first_name, ud.father_name, ud.faculty_id, ud.department_id, ud.role, ud.educational_direction, g.name AS group_name
		FROM user_data ud LEFT JOIN "group" g ON ud.id = g.user_id
		WHERE ud.university_id = $1 AND ud.is_deleted = false AND ud.role = $2;
	`

	rows, err := r.db.QueryContext(ctx, query, universityID, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userData []entities.GetUserDataResponse
	for rows.Next() {
		var user entities.GetUserDataResponse
		err = rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.FatherName, &user.FacultyID, &user.DepartmentID, &user.Role, &user.EducationalDirection, &user.Group)
		if err != nil {
			return nil, err
		}
		userData = append(userData, user)
	}
	return &userData, nil
}

func (r *UserDataRepository) GetEducationalDirection(ctx context.Context, userId int) (string, error) {
	var educationalDirection string
	query := `SELECT educational_direction FROM user_data WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, userId).Scan(&educationalDirection)
	if err != nil {
		return "", err
	}

	return educationalDirection, nil
}

func (r *UserDataRepository) GetById(ctx context.Context, userID int) (*entities.UserData, error) {
	var user entities.UserData

	query := `
		SELECT first_name, last_name, father_name, role, faculty_id, department_id, educational_direction, university_id
		FROM user_data 
		WHERE id = $1 LIMIT 1;
	`

	err := r.db.QueryRowContext(ctx, query, userID).Scan(&user.FirstName, &user.LastName, &user.FatherName, &user.Role, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection, &user.UniversityID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
