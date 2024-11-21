package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"server/internal/entities"
)

type UserScheduleRepository struct {
	db *sql.DB
}

func NewMyScheduleRepository(db *sql.DB) *UserScheduleRepository {
	return &UserScheduleRepository{
		db: db,
	}
}

func (r *UserScheduleRepository) Create(ctx context.Context, userSchedule *entities.UserSchedule) (int, error) {
	if userSchedule.Name == "" || userSchedule.Date == "" || userSchedule.TimeStart == "" || userSchedule.TimeEnd == "" {
		return 0, errors.New("one of the fields of the user schedule is empty")
	}

	query := `SELECT * FROM my_schedule WHERE date=$1 AND time_start=$2`
	row := r.db.QueryRowContext(ctx, query, userSchedule.Date, userSchedule.TimeStart)
	var tmp interface{}
	err := row.Scan(&tmp)
	if !(err == sql.ErrNoRows) {
		return 0, errors.New("my_schedule with these fields already exists")
	}

	var id int
	query = `INSERT INTO my_schedule (user_data_id, name, date, time_start, time_end) VALUES ($1) RETURNING id`

	err = r.db.QueryRowContext(ctx, query, userSchedule.UserDataId, userSchedule.Name, userSchedule.Date, userSchedule.TimeStart, userSchedule.TimeEnd).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println(err)
	return id, nil
}

func (r *UserScheduleRepository) GetByUserId(ctx context.Context, userId int) (*entities.UserSchedule, error) {
	var mySchedule entities.UserSchedule
	query := `SELECT * FROM my_schedule WHERE user_data_id = $1`
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&mySchedule.Id, &mySchedule.Name, &mySchedule.Date, &mySchedule.TimeStart, &mySchedule.TimeEnd)
	if err != nil {
		return nil, err
	}
	return &mySchedule, nil
}

func (r *UserScheduleRepository) Update(ctx context.Context, mySchedule *entities.UserSchedule) error {
	query := `UPDATE my_schedule SET name = $1, date = $2, time_start = $3 , time_end = $4 WHERE id = $5`
	err := r.db.QueryRowContext(ctx, query, mySchedule.Name, mySchedule.Date, mySchedule.TimeStart, mySchedule.TimeEnd, mySchedule.Id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserScheduleRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM my_schedule WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Err()
	if err != nil {
		return err
	}
	return nil
}
