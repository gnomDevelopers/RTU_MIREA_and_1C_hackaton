package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"server/internal/entities"
)

type ClassRepository struct {
	db *sql.DB
}

func NewClassRepository(db *sql.DB) *ClassRepository {
	return &ClassRepository{
		db: db,
	}
}

func (r *ClassRepository) Create(ctx context.Context, classes *[]entities.Class) ([]int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return []int{}, errors.New("failed to begin transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var ids []int
	for _, class := range *classes {
		if class.Name == "" || class.Type == "" || class.Auditory == "" || class.Date == "" || class.Week == 0 || class.Weekday == 0 || class.TimeStart == "" || class.TimeEnd == "" {
			return []int{}, errors.New("one of the fields of the class is empty")
		}

		var id int
		var exists bool
		checkQuery := `SELECT EXISTS (
                    SELECT 1 FROM class 
                    WHERE date = $1 AND weekday = $2 AND week = $3 AND time_start = $4 AND time_end = $5
                );`

		err = tx.QueryRowContext(ctx, checkQuery, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd).Scan(&exists)
		if err != nil {
			return []int{}, err
		}

		// Если запись не существует, выполняем вставку
		if !exists {
			insertQuery := `INSERT INTO class (name, group_names, teacher_names, type, auditory, date, weekday, week, time_start, time_end) 
                    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
                    RETURNING id;`

			err = tx.QueryRowContext(ctx, insertQuery, class.Name, pq.Array(class.GroupNames), pq.Array(class.TeacherNames), class.Type, class.Auditory, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd).Scan(&id)
			if err != nil {
				return []int{}, err
			}
			ids = append(ids, id)
		} else {
			return []int{}, errors.New("class with such fields already exists")
		}
	}
	err = tx.Commit()
	if err != nil {
		return []int{}, errors.New("problem with transaction commit")
	}

	return ids, nil
}

func (r *ClassRepository) GetById(ctx context.Context, id int) (*entities.Class, error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}

	var class entities.Class
	query := `SELECT * FROM class WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&class.Id, &class.Name, &class.GroupNames, &class.TeacherNames, &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd)
	if err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *ClassRepository) GetByGroupName(ctx context.Context, groupName string) (*[]entities.Class, error) {
	if groupName == "" {
		return nil, errors.New("groupName is empty")
	}

	var classes []entities.Class
	query := `SELECT * FROM class WHERE $1=ANY(group_names)`
	rows, err := r.db.QueryContext(ctx, query, groupName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, &class.GroupNames, &class.TeacherNames, &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return &classes, nil
}

func (r *ClassRepository) GetByTeacherName(ctx context.Context, teacherName string) (*[]entities.Class, error) {
	if teacherName == "" {
		return nil, errors.New("")
	}

	var classes []entities.Class
	query := `SELECT * FROM class WHERE $1=ANY(teacher_names)`
	rows, err := r.db.QueryContext(ctx, query, teacherName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, &class.GroupNames, &class.TeacherNames, &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return &classes, nil
}

func (r *ClassRepository) GetByAuditoryId(ctx context.Context, auditoryId int) (*[]entities.Class, error) {
	if auditoryId == 0 {
		return nil, errors.New("")
	}

	var classes []entities.Class
	query := `SELECT * FROM class WHERE auditory=$1`
	rows, err := r.db.QueryContext(ctx, query, auditoryId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, &class.GroupNames, &class.TeacherNames, &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return &classes, nil
}

func (r *ClassRepository) Update(ctx context.Context, class *entities.Class) error {
	if class.Name == "" || class.Type == "" || class.Auditory == "" || class.Date == "" || class.Week == 0 || class.Weekday == 0 || class.TimeStart == "" || class.TimeEnd == "" {
		return errors.New("")
	}

	query := `UPDATE class SET name=$1, group_names=$2, teacher_names=$3, type=$4, auditory=$5, date=$6, weekday=$7, week=$8, time_start=$9, time_end=$10 WHERE id=$11`

	_, err := r.db.ExecContext(ctx, query, class.Name, class.GroupNames, class.TeacherNames, class.Type, class.Auditory, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd, class.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ClassRepository) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("")
	}

	query := `DELETE FROM class WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
