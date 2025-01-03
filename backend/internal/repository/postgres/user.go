package postgres

import (
	"context"
	"database/sql"
	"errors"
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
	query := "SELECT password, id, email, last_name, first_name, father_name, university_id, role, faculty_id, department_id, educational_direction FROM users WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Password, &user.ID, &user.Email, &user.LastName, &user.FirstName, &user.FatherName, &user.UniversityID, &user.Role, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil

}

func (r *UserRepository) GetInfoById(ctx context.Context, id int) (*entities.UserInfo, error) {
	user := entities.UserInfo{}

	query := `
		SELECT 
			u.id, 
			u.email, 
			u.last_name, 
			u.first_name, 
			u.father_name, 
			un.name AS university_name, 
			u.role, 
			u.faculty_id, 
			u.department_id, 
			u.educational_direction,
			u.group_id
		FROM users u
		LEFT JOIN university un ON u.university_id = un.id
		WHERE u.id = $1 AND u.is_deleted = false
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.LastName,
		&user.FirstName,
		&user.FatherName,
		&user.UniversityName,
		&user.Role,
		&user.FacultyID,
		&user.DepartmentID,
		&user.EducationalDirection,
		&user.GroupID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, login string) (*entities.User, error) {
	user := entities.User{}
	query := "SELECT password, id, email, last_name, first_name, father_name, university_id, role, faculty_id, department_id, educational_direction FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, login).Scan(&user.Password, &user.ID, &user.Email, &user.LastName, &user.FirstName, &user.FatherName, &user.UniversityID, &user.Role, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection)
	if err != nil {
		return &entities.User{}, err
	}

	return &user, nil

}

func (r *UserRepository) GetByUniversity(ctx context.Context, university string) (*[]entities.User, error) {
	if university == "" {
		return nil, errors.New("university is empty")
	}

	var users []entities.User

	query := `SELECT users.id, email, last_name, first_name, father_name, university_id, role, group_id, faculty_id, department_id, educational_direction FROM users JOIN university ON users.university_id = university.id WHERE university.name = $1;`
	rows, err := r.db.QueryContext(ctx, query, university)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.ID, &user.Email, &user.LastName, &user.FirstName, &user.FatherName, &user.UniversityID, &user.Role, &user.GroupID, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("there are no users with such parameters")
	}

	return &users, nil
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

func (r *UserRepository) UpdateRole(ctx context.Context, userID int, newRole string) error {
	query := `UPDATE users SET role = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, newRole, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAllTeachers(ctx context.Context, university string) (*[]entities.User, error) {
	var users []entities.User

	query := "SELECT password, id, email, last_name, first_name, father_name, university_id, role, faculty_id, department_id, educational_direction FROM users JOIN university ON users.university_id = university.id WHERE role = $1 AND university.name = $2"
	rows, err := r.db.QueryContext(ctx, query, "Преподаватель")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.Password, &user.ID, &user.Email, &user.LastName, &user.FirstName, &user.FatherName, &user.UniversityID, &user.Role, &user.FacultyID, &user.DepartmentID, &user.EducationalDirection)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
