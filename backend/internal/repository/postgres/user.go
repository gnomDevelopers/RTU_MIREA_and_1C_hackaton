package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetById(ctx context.Context, id int) (*entities.User, error) {
	user := entities.User{}
	query := "SELECT id, email, last_name, first_name, father_name, university_id, role, faculty_id, department_id, educational_direction FROM users WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.LastName, &user.FirstName, &user.FatherName, &user.UniversityID, &user.Role, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil

}

func (r *UserRepository) GetByEmail(ctx context.Context, login string) (*entities.User, error) {
	user := entities.User{}
	query := "SELECT id, email, last_name, first_name, father_name, university_id, role, faculty_id, department_id, educational_direction FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Email, &user.LastName, &user.FirstName, &user.FatherName, &user.UniversityID, &user.Role, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil

}

func (r *UserRepository) Exists(ctx context.Context, email string) (bool, error) {
	exists := 0
	query := "SELECT 1 FROM users WHERE email = $1 LIMIT 1"

	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

//createAdmin := &entities.User{
//		Email:                req.Email,
//		Password:             req.Password,
//		FirstName:            req.FirstName,
//		LastName:             req.LastName,
//		FatherName:           req.FatherName,
//		Role:                 "Администратор",
//		UniversityID:         1,
//		FacultyID:            1,
//		GroupID:                1,
//		DepartmentID:         1,
//		EducationalDirection: "admin",
//	}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var lastInsertId int
	query := `
		INSERT INTO users (email, password, first_name, last_name, father_name, role, university_id, faculty_id, group_id, department_id, educational_direction)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password, user.FirstName, user.LastName, user.FatherName, user.Role, user.UniversityID, user.FacultyID, user.GroupID, user.DepartmentID, user.EducationalDirection).Scan(&lastInsertId)
	if err != nil {
		return &entities.User{}, err
	}
	user.ID = lastInsertId
	return user, nil
}
