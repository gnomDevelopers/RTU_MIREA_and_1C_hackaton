package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Exists(ctx context.Context, name string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM "group" WHERE name = $1)`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&exists)
	return exists, err
}

func (r *GroupRepository) Create(ctx context.Context, group *entities.CreateGroupRequest) (int, error) {
	var id int
	query := `INSERT INTO "group" (name, user_id) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, group.Name, group.UserID).Scan(&id)
	return id, err
}

func (r *GroupRepository) GetById(ctx context.Context, id int) (*entities.Group, error) {
	var group entities.Group
	query := `SELECT id, name, user_id FROM "group" WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&group.ID, &group.Name, &group.UserID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &group, err
}

func (r *GroupRepository) GetByName(ctx context.Context, name string) (*entities.Group, error) {
	var group entities.Group
	query := `SELECT id, name, user_id FROM "group" WHERE name = $1`
	row := r.db.QueryRowContext(ctx, query, name)
	err := row.Scan(&group.ID, &group.Name, &group.UserID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &group, err
}

func (r *GroupRepository) GetAll(ctx context.Context) (*[]entities.Group, error) {
	query := `SELECT id, name, user_id FROM "group"`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []entities.Group
	for rows.Next() {
		var group entities.Group
		if err := rows.Scan(&group.ID, &group.Name, &group.UserID); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return &groups, rows.Err()
}