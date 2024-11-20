package postgres

import (
	"database/sql"
)

type UserDataRepository struct {
	db *sql.DB
}

func NewUserDataRepository(db *sql.DB) *UserDataRepository {
	return &UserDataRepository{
		db: db,
	}
}

//func (r *UserDataRepository) Add(ctx context.Context, userData *entities.UserData) error {
//	tx, err := r.db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	defer tx.Rollback()
//
//	var universityID string
//	query := `SELECT id FROM university WHERE name = $1`
//	row := tx.QueryRowContext(ctx, query, userData.University)
//	err = row.Scan(&universityID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//
//			return fmt.Errorf("university with name '%s' not found", userData.University)
//		}
//		return err
//	}
//
//	var facultyID string
//	query = `SELECT id FROM faculty WHERE name = $1`
//	row = tx.QueryRowContext(ctx, query, userData.Faculty)
//	err = row.Scan(&facultyID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return fmt.Errorf("faculty with name '%s' not found", userData.University)
//		}
//		tx.Rollback()
//		return err
//	}
//
//	var departmentID string
//	query = `SELECT id FROM depatment WHERE name = $1`
//	row = tx.QueryRowContext(ctx, query, userData.Department)
//	err = row.Scan(&departmentID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return fmt.Errorf("department with name '%s' not found", userData.University)
//		}
//		tx.Rollback()
//		return err
//	}
//
//	query := `
//        INSERT INTO user_data (
//            id, last_name, first_name, father_name, university_id,
//            permission_id, faculty_id, department_id, is_deleted,
//            is_password_changed, created_at
//        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
//    `
//
//}
