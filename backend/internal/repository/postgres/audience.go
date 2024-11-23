package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type AudienceRepository struct {
	db *sql.DB
}

func NewAudienceRepository(db *sql.DB) *AudienceRepository {
	return &AudienceRepository{
		db: db,
	}
}

func (r *AudienceRepository) Create(ctx context.Context, audiences *[]entities.Audience) ([]int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return []int{}, errors.New("")
	}

	var ids []int
	for _, audience := range *audiences {
		if audience.Name == "" || audience.Campus == "" || audience.Type == "" || audience.Profile == "" || audience.Capacity == 0 {
			return []int{}, errors.New("")
		}

		query := `SELECT * FROM auditory WHERE name=$1 AND campus=$2 AND type=$3 AND profile=$4 AND capacity=$5`
		row := tx.QueryRowContext(ctx, query, audience.Name, audience.Campus, audience.Type, audience.Profile, audience.Capacity)
		var tmp interface{}
		err := row.Scan(&tmp)
		if !(err == sql.ErrNoRows) {
			return []int{}, errors.New("audience with these fields already exists")
		}

		var id int
		query = `INSERT INTO auditory (name, campus, type, profile, capacity) VALUES($1, $2, $3, $4, $5) RETURNING id`

		err = tx.QueryRowContext(ctx, query, audience.Name, audience.Campus, audience.Type, audience.Profile, audience.Capacity).Scan(&id)
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

func (r *AudienceRepository) GetById(ctx context.Context, id int) (*entities.Audience, error) {
	if id == 0 {
		return nil, errors.New("")
	}

	var audience entities.Audience
	query := `SELECT * FROM auditory WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&audience.Id, &audience.Name, &audience.Campus, &audience.Type, &audience.Profile, &audience.Capacity)
	if err != nil {
		return nil, err
	}
	return &audience, nil
}

func (r *AudienceRepository) GetByCampus(ctx context.Context, campus string) (*[]entities.Audience, error) {
	if campus == "" {
		return nil, errors.New("")
	}

	var audiences []entities.Audience
	query := `SELECT * FROM auditory WHERE campus = $1`
	rows, err := r.db.QueryContext(ctx, query, campus)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var audience entities.Audience
		err = rows.Scan(&audience.Id, &audience.Name, &audience.Campus, &audience.Type, &audience.Profile, &audience.Capacity)
		if err != nil {
			return nil, err
		}
		audiences = append(audiences, audience)
	}

	if len(audiences) == 0 {
		return nil, errors.New("there are no audiences with such parameters")
	}

	return &audiences, nil
}

func (r *AudienceRepository) GetByType(ctx context.Context, typeStr string) (*[]entities.Audience, error) {
	if typeStr == "" {
		return nil, errors.New("")
	}

	var audiences []entities.Audience
	query := `SELECT * FROM auditory WHERE type = $1`
	rows, err := r.db.QueryContext(ctx, query, typeStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var audience entities.Audience
		err = rows.Scan(&audience.Id, &audience.Name, &audience.Campus, &audience.Type, &audience.Profile, &audience.Capacity)
		if err != nil {
			return nil, err
		}
		audiences = append(audiences, audience)
	}

	if len(audiences) == 0 {
		return nil, errors.New("there are no audiences with such parameters")
	}

	return &audiences, nil
}

func (r *AudienceRepository) GetByProfile(ctx context.Context, profile string) (*[]entities.Audience, error) {
	if profile == "" {
		return nil, errors.New("")
	}

	var audiences []entities.Audience
	query := `SELECT * FROM auditory WHERE profile = $1`
	rows, err := r.db.QueryContext(ctx, query, profile)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var audience entities.Audience
		err = rows.Scan(&audience.Id, &audience.Name, &audience.Campus, &audience.Type, &audience.Profile, &audience.Capacity)
		if err != nil {
			return nil, err
		}
		audiences = append(audiences, audience)
	}

	if len(audiences) == 0 {
		return nil, errors.New("there are no audiences with such parameters")
	}

	return &audiences, nil
}

func (r *AudienceRepository) GetByCapacity(ctx context.Context, capacity int) (*[]entities.Audience, error) {
	if capacity == 0 {
		return nil, errors.New("")
	}

	var audiences []entities.Audience
	query := `SELECT * FROM auditory WHERE capacity >= $1`
	rows, err := r.db.QueryContext(ctx, query, capacity)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var audience entities.Audience
		err = rows.Scan(&audience.Id, &audience.Name, &audience.Campus, &audience.Type, &audience.Profile, &audience.Capacity)
		if err != nil {
			return nil, err
		}
		audiences = append(audiences, audience)
	}

	if len(audiences) == 0 {
		return nil, errors.New("there are no audiences with such parameters")
	}

	return &audiences, nil
}

func (r *AudienceRepository) Update(ctx context.Context, audience *entities.Audience) error {
	if audience.Name == "" || audience.Campus == "" || audience.Type == "" || audience.Profile == "" || audience.Capacity == 0 {
		return errors.New("")
	}

	query := `UPDATE auditory SET name=$1, campus=$2, type=$3, profile=$4, capacity=$5 WHERE id=$6`

	_, err := r.db.ExecContext(ctx, query, audience.Name, audience.Campus, audience.Type, audience.Profile, audience.Capacity, audience.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AudienceRepository) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("")
	}

	query := `DELETE FROM auditory WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
