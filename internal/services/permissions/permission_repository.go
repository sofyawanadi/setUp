package repository

import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

type Permissions struct {
	ID   int64
	Name string
}

type PermissionsRepository interface {
	Create(ctx context.Context, ug *Permissions) error
	GetByID(ctx context.Context, id int64) (*Permissions, error) 
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]Permissions, error) 
	Update(ctx context.Context, ug *Permissions) error 
}

type permissionsRepository struct {
	DB *sql.DB
	log *zap.Logger
}

func NewPermissionsRepository(db *sql.DB, log *zap.Logger) PermissionsRepository {
	return &permissionsRepository{db,log}
}

// Create
func (r *permissionsRepository) Create(ctx context.Context, ug *Permissions) error {
	query := "INSERT INTO permissions (name) VALUES (?)"
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
func (r *permissionsRepository) GetByID(ctx context.Context, id int64) (*Permissions, error) {
	query := "SELECT id, name FROM permissions WHERE id = ?"
	row := r.DB.QueryRowContext(ctx, query, id)
	var ug Permissions
	if err := row.Scan(&ug.ID, &ug.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ug, nil
}

// Read (Get All)
func (r *permissionsRepository) GetAll(ctx context.Context) ([]Permissions, error) {
	query := "SELECT id, name FROM permissions"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Permissions
	for rows.Next() {
		var ug Permissions
		if err := rows.Scan(&ug.ID, &ug.Name); err != nil {
			return nil, err
		}
		groups = append(groups, ug)
	}
	return groups, nil
}

// Update
func (r *permissionsRepository) Update(ctx context.Context, ug *Permissions) error {
	query := "UPDATE permissions SET name = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, ug.Name, ug.ID)
	return err
}

// Delete
func (r *permissionsRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM permissions WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}