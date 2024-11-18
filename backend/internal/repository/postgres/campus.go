package postgres

import (
	"context"
	"database/sql"
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

func (c *CampusRepository) Create(ctx context.Context, universityId int, name, address string) (int, error) {
	var id int
	query := `INSERT INTO campus (name, university_id, address) VALUES ($1, $2, $3) RETURNING id`

	err := c.db.QueryRowContext(ctx, query, name, universityId, address).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *CampusRepository) GetById(ctx context.Context, id int) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE id = $1`
	err := c.db.QueryRowContext(ctx, query, id).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (c *CampusRepository) GetByAddress(ctx context.Context, address string) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE address = $1`
	err := c.db.QueryRowContext(ctx, query, address).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (c *CampusRepository) GetByName(ctx context.Context, name string) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE name = $1`
	err := c.db.QueryRowContext(ctx, query, name).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (c *CampusRepository) GetByUniversityId(ctx context.Context, universityId int) (*entities.Campus, error) {
	var campus entities.Campus
	query := `SELECT * FROM campus WHERE university_id = $1`
	err := c.db.QueryRowContext(ctx, query, universityId).Scan(&campus.Id, &campus.Name, &campus.UniversityId, &campus.Address)
	if err != nil {
		return nil, err
	}
	return &campus, nil
}

func (c *CampusRepository) GetAll(ctx context.Context) (*[]entities.Campus, error) {
	var campuses []entities.Campus
	query := `SELECT * FROM campus`
	rows, err := c.db.QueryContext(ctx, query)
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

func (c *CampusRepository) Update(ctx context.Context, id, universityId int, name, address string) error {
	query := `UPDATE campus SET name = $1, university_id = $2, address = $3 WHERE id = $4`
	err := c.db.QueryRowContext(ctx, query, name, universityId, address, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *CampusRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM campus WHERE id = $1`
	err := c.db.QueryRowContext(ctx, query, id).Err()
	if err != nil {
		return err
	}
	return nil
}
