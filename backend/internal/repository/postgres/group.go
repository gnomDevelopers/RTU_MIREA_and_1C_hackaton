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
	query := `INSERT INTO "group" (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, group.Name).Scan(&id)
	return id, err
}

func (r *GroupRepository) GetById(ctx context.Context, id int) (*entities.Group, error) {
	var group entities.Group
	query := `SELECT id, name FROM "group" WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&group.ID, &group.Name)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &group, err
}

func (r *GroupRepository) GetByUserID(ctx context.Context, userID int) (*entities.Group, error) {
	group := &entities.Group{}

	query := `
		SELECT g.id, g.name
		FROM users u JOIN "group" g ON u.group_id = g.id
		WHERE u.id = $1 LIMIT 1; 
`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepository) GetByName(ctx context.Context, name string) (*entities.Group, error) {
	var group entities.Group
	query := `SELECT id, name, FROM "group" WHERE name = $1`
	row := r.db.QueryRowContext(ctx, query, name)
	err := row.Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepository) GetAll(ctx context.Context) (*[]entities.Group, error) {
	query := `SELECT id, name FROM "group"`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []entities.Group
	for rows.Next() {
		var group entities.Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return &groups, nil
}

func (r *GroupRepository) GetGroupMembers(ctx context.Context, group string) (*[]entities.GroupMember, error) {
	query := `
		SELECT users.id AS user_id, users.last_name, users.first_name, users.father_name,
    "group".name AS group_name
	FROM users
	JOIN "group" ON users.group_id = "group".id
	WHERE "group".name = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var members []entities.GroupMember
	for rows.Next() {
		var member entities.GroupMember
		if err := rows.Scan(&member.ID, &member.LastName, &member.FirstName, &member.FatherName); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return &members, nil
}
