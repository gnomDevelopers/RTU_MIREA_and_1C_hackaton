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

		query := `SELECT 1 FROM class WHERE auditory_id=$1 AND date=$2 AND weekday=$3 AND week=$4 AND time_start=$5 AND time_end=$6`
		row := tx.QueryRowContext(ctx, query, class.Auditory, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd)
		if err := row.Scan(new(interface{})); err != nil && err != sql.ErrNoRows {
			return []int{}, err
		}
		if err == nil {
			return []int{}, errors.New("class with these fields already exists")
		}

		var id int
		query = `INSERT INTO class (name, academic_discipline_id, group_names, teacher_names, type, auditory_id, date, weekday, week, time_start, time_end) 
				 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`

		err = tx.QueryRowContext(ctx, query, class.Name, pq.Array(class.GroupNames), pq.Array(class.TeacherNames), class.Type, class.Auditory, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd).Scan(&id)
		if err != nil {
			return []int{}, err
		}
		ids = append(ids, id)
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
	query := `SELECT * FROM class WHERE auditory_id=$1`
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

	query := `UPDATE class SET name=$1, academic_discipline_id=$2, group_names=$3, teacher_names=$4, type=$5, auditory_id=$6, date=$7, weekday=$8, week=$9, time_start=$10, time_end=$11 WHERE id=$12`

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
