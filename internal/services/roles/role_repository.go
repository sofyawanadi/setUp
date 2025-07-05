package repository

import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

type Role struct {
	ID   int64
	Name string
}

type RoleRepository interface {
	Create(ctx context.Context, ug *Role) error
	GetByID(ctx context.Context, id int64) (*Role, error) 
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]Role, error) 
	Update(ctx context.Context, ug *Role) error 
}

type roleRepository struct {
	DB *sql.DB
	log *zap.Logger
}

func NewRoleRepository(db *sql.DB, log *zap.Logger) RoleRepository {
	return &roleRepository{db,log}
}

// Create
func (r *roleRepository) Create(ctx context.Context, ug *Role) error {
	query := "INSERT INTO roles (name) VALUES (?)"
	result, err := r.DB.ExecContext(ctx, query, ug.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ug.ID = id
	return nil
}

// Read (Get by ID)
func (r *roleRepository) GetByID(ctx context.Context, id int64) (*Role, error) {
	query := "SELECT id, name FROM roles WHERE id = ?"
	row := r.DB.QueryRowContext(ctx, query, id)
	var ug Role
	if err := row.Scan(&ug.ID, &ug.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ug, nil
}

// Read (Get All)
func (r *roleRepository) GetAll(ctx context.Context) ([]Role, error) {
	query := "SELECT id, name FROM roles"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Role
	for rows.Next() {
		var ug Role
		if err := rows.Scan(&ug.ID, &ug.Name); err != nil {
			return nil, err
		}
		groups = append(groups, ug)
	}
	return groups, nil
}

// Update
func (r *roleRepository) Update(ctx context.Context, ug *Role) error {
	query := "UPDATE roles SET name = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, ug.Name, ug.ID)
	return err
}

// Delete
func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM roles WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}