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
                    WHERE date = $1 AND weekday = $2 AND week = $3 AND time_start = $4 AND time_end = $5 AND auditory = $6
                );`
		err = tx.QueryRowContext(ctx, checkQuery, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd, class.Auditory).Scan(&exists)
		if err != nil && err != sql.ErrNoRows {
			return []int{}, err
		}

		// Если запись не существует, выполняем вставку
		if !exists {
			insertQuery := `INSERT INTO class (name, group_names, teacher_names, type, auditory, date, weekday, week, time_start, time_end, university) 
                    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
                    RETURNING id;`

			err = tx.QueryRowContext(ctx, insertQuery, class.Name, pq.Array(class.GroupNames), pq.Array(class.TeacherNames), class.Type, class.Auditory, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd, class.UniversityStr).Scan(&id)
			if err != nil {
				return []int{}, err
			}
			ids = append(ids, id)
		} else {
			if class.Type == "ЛК" || class.Auditory == "Дистанционно" {
				continue
			}
			err = tx.Rollback()
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
	err := r.db.QueryRowContext(ctx, query, id).Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd, &class.UniversityStr)
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

	dateLimit := "09-08-24"
	query := `SELECT * FROM class WHERE $1=ANY(group_names) AND date < $2`
	rows, err := r.db.QueryContext(ctx, query, groupName, dateLimit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd, &class.UniversityStr)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return nil, errors.New("there are no classes with such parameters")
	}

	return &classes, nil
}

func (r *ClassRepository) GetByTeacherName(ctx context.Context, teacherName, university string) (*[]entities.Class, error) {
	if teacherName == "" {
		return nil, errors.New("")
	}

	var classes []entities.Class
	dateLimit := "09-08-24"
	query := `SELECT * FROM class WHERE $1=ANY(teacher_names) AND date < $2 AND university = $3`
	rows, err := r.db.QueryContext(ctx, query, teacherName, dateLimit, university)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd, &class.UniversityStr)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return nil, errors.New("there are no classes with such parameters")
	}

	return &classes, nil
}

func (r *ClassRepository) GetByName(ctx context.Context, name, university string) (*[]entities.Class, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	var classes []entities.Class
	dateLimit := "09-08-24"
	query := `SELECT * FROM class WHERE $1=name AND type='ФАКУЛЬТАТИВ' AND date < $2 AND university = $3;`
	rows, err := r.db.QueryContext(ctx, query, name, dateLimit, university)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd, &class.UniversityStr)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return nil, errors.New("there are no classes with such parameters")
	}

	return &classes, nil
}

func (r *ClassRepository) GetByNameAndGroup(ctx context.Context, name, group string) (*[]entities.Class, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	var classes []entities.Class

	query := `SELECT * FROM class WHERE $1=name AND $2=ANY(group_names);`
	rows, err := r.db.QueryContext(ctx, query, name, group)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd, &class.UniversityStr)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return nil, errors.New("there are no classes with such parameters")
	}

	return &classes, nil
}
func (r *ClassRepository) GetByNameAndGroupWithoutLk(ctx context.Context, name, group string) (*[]entities.Class, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	var classes []entities.Class

	query := `SELECT * FROM class WHERE $1=name AND $2=ANY(group_names) AND type<>'ЛК';`
	rows, err := r.db.QueryContext(ctx, query, name, group)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd, &class.UniversityStr)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return nil, errors.New("there are no classes with such parameters")
	}

	return &classes, nil
}

func (r *ClassRepository) GetByAuditory(ctx context.Context, auditory string) (*[]entities.Class, error) {
	if auditory == "" {
		return nil, errors.New("")
	}

	var classes []entities.Class
	dateLimit := "09-08-24"
	query := `SELECT * FROM class WHERE auditory=$1 AND date < $2;`
	rows, err := r.db.QueryContext(ctx, query, auditory, dateLimit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class entities.Class
		err = rows.Scan(&class.Id, &class.Name, pq.Array(&class.GroupNames), pq.Array(&class.TeacherNames), &class.Type, &class.Auditory, &class.Date, &class.Weekday, &class.Week, &class.TimeStart, &class.TimeEnd)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return nil, errors.New("there are no classes with such parameters")
	}

	return &classes, nil
}

func (r *ClassRepository) SearchGroups(ctx context.Context, university string) ([]string, error) {
	query := `
		SELECT unnest(group_names)
		FROM class
		WHERE university = $1
	`

	rows, err := r.db.QueryContext(ctx, query, university)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groupMap := make(map[string]struct{})
	for rows.Next() {
		var groupName string
		if err := rows.Scan(&groupName); err != nil {
			return nil, err
		}
		groupMap[groupName] = struct{}{}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	uniqueGroupNames := make([]string, 0, len(groupMap))
	for groupName := range groupMap {
		uniqueGroupNames = append(uniqueGroupNames, groupName)
	}

	if len(uniqueGroupNames) == 0 {
		return nil, errors.New("there are no groups")
	}

	return uniqueGroupNames, nil
}

func (r *ClassRepository) SearchTeachers(ctx context.Context, university string) ([]string, error) {
	query := `
		SELECT unnest(teacher_names)
		FROM class
		WHERE university = $1
	`

	rows, err := r.db.QueryContext(ctx, query, university)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teacherMap := make(map[string]struct{})
	for rows.Next() {
		var teacherName string
		if err := rows.Scan(&teacherName); err != nil {
			return nil, err
		}
		teacherMap[teacherName] = struct{}{}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	uniqueTeacherNames := make([]string, 0, len(teacherMap))
	for groupName := range teacherMap {
		uniqueTeacherNames = append(uniqueTeacherNames, groupName)
	}

	return uniqueTeacherNames, nil
}

func (r *ClassRepository) SearchNames(ctx context.Context, university string) ([]string, error) {
	query := `
		SELECT name
		FROM class
		WHERE university = $1 AND type='ФАКУЛЬТАТИВ'
	`

	rows, err := r.db.QueryContext(ctx, query, university)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nameMap := make(map[string]struct{})
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		nameMap[name] = struct{}{}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	uniqueNames := make([]string, 0, len(nameMap))
	for name := range nameMap {
		uniqueNames = append(uniqueNames, name)
	}

	return uniqueNames, nil
}

func (r *ClassRepository) SearchNamesWithGroup(ctx context.Context, group string) ([]string, error) {
	query := `
		SELECT name
		FROM class
		WHERE $1=ANY(group_names)
	`

	rows, err := r.db.QueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nameMap := make(map[string]struct{})
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		nameMap[name] = struct{}{}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	uniqueNames := make([]string, 0, len(nameMap))
	for name := range nameMap {
		uniqueNames = append(uniqueNames, name)
	}

	return uniqueNames, nil
}

func (r *ClassRepository) Update(ctx context.Context, class *entities.Class) error {
	if class.Name == "" || class.Type == "" || class.Auditory == "" || class.Date == "" || class.Week == 0 || class.Weekday == 0 || class.TimeStart == "" || class.TimeEnd == "" {
		return errors.New("")
	}

	query := `UPDATE class SET name=$1, group_names=$2, teacher_names=$3, type=$4, auditory=$5, date=$6, weekday=$7, week=$8, time_start=$9, time_end=$10 WHERE id=$11`

	_, err := r.db.ExecContext(ctx, query, class.Name, pq.Array(class.GroupNames), pq.Array(class.TeacherNames), class.Type, class.Auditory, class.Date, class.Weekday, class.Week, class.TimeStart, class.TimeEnd, class.Id)
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
