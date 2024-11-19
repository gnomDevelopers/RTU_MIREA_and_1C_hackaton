package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type CampusRepository struct {
	db *sql.DB
}

func NewCampusRepository(db *sql.DB) *CampusRepository {
	return &CampusRepository{
		db: db,
	}
}

func (r *CampusRepository) Create(ctx context.Context, campuses *[]entities.Campus) ([]int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return []int{}, errors.New("")
	}

	var ids []int
	for _, campus := range *campuses {
		if campus.Name == "" || campus.Address == "" {
			return []int{}, errors.New("")
		}

		query := `SELECT * FROM campus WHERE name=$1 AND address=$2`
		row := tx.QueryRowContext(ctx, query, campus.Name, campus.Address)
		var tmp interface{}
		err := row.Scan(&tmp)
		if !(err == sql.ErrNoRows) {
			return []int{}, errors.New("class with these fields already exists")
		}

		var id int
		query = `INSERT INTO campus (name, university_id, address) VALUES ($1, $2, $3) RETURNING id`

		err = tx.QueryRowContext(ctx, query, campus.Name, campus.UniversityId, campus.Address).Scan(&id)
		if err != nil {
			tx.Rollback()
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

func (r *CampusRepository) GetById(ctx context.Context, id int) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (r *CampusRepository) GetByAddress(ctx context.Context, address string) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE address = $1`
	err := r.db.QueryRowContext(ctx, query, address).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (r *CampusRepository) GetByName(ctx context.Context, name string) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE name = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (r *CampusRepository) GetByUniversityId(ctx context.Context, universityId int) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE university_id = $1`
	err := r.db.QueryRowContext(ctx, query, universityId).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (r *CampusRepository) GetAll(ctx context.Context) (*[]entities.Campus, error) {
	var campuses []entities.Campus
	query := `SELECT * FROM campus`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var campus entities.Campus
		err = rows.Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
		if err != nil {
			return nil, err
		}
		campuses = append(campuses, campus)
	}

	return &campuses, nil
}

func (r *CampusRepository) Update(ctx context.Context, campus *entities.Campus) error {
	query := `UPDATE campus SET name = $1, university_id = $2, address = $3 WHERE id = $4`
	err := r.db.QueryRowContext(ctx, query, campus.Name, campus.UniversityId, campus.Address, campus.Id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *CampusRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM campus WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Err()
	if err != nil {
		return err
	}
	return nil
}
