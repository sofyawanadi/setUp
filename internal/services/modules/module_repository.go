package services

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type ModuleRepository interface {
	Create(ctx context.Context, ug *Module) error
	GetByID(ctx context.Context, id int64) (*Module, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]Module, error)
	Update(ctx context.Context, ug *Module) error
}

type moduleRepository struct {
	DB  *sql.DB
	log *zap.Logger
}

func NewModuleRepository(db *sql.DB, log *zap.Logger) ModuleRepository {
	return &moduleRepository{db, log}
}

// Create
func (r *moduleRepository) Create(ctx context.Context, ug *Module) error {
	query := "INSERT INTO modules (name) VALUES (?)"
	_, err := r.DB.ExecContext(ctx, query, ug.Name)
	if err != nil {
		return err
	}
	return nil
}

// Read (Get by ID)
func (r *moduleRepository) GetByID(ctx context.Context, id int64) (*Module, error) {
	query := "SELECT id, name FROM modules WHERE id = ?"
	row := r.DB.QueryRowContext(ctx, query, id)
	var ug Module
	if err := row.Scan(&ug.ID, &ug.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ug, nil
}

// Read (Get All)
func (r *moduleRepository) GetAll(ctx context.Context) ([]Module, error) {
	query := "SELECT id, name FROM modules"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Module
	for rows.Next() {
		var ug Module
		if err := rows.Scan(&ug.ID, &ug.Name); err != nil {
			return nil, err
		}
		groups = append(groups, ug)
	}
	return groups, nil
}

// Update
func (r *moduleRepository) Update(ctx context.Context, ug *Module) error {
	query := "UPDATE modules SET name = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, ug.Name, ug.ID)
	return err
}

// Delete
func (r *moduleRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM modules WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
