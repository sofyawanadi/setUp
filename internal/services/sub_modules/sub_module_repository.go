package services
import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

type SubModuleRepository interface {
	Create(ctx context.Context, ug *SubModules) error
	GetByID(ctx context.Context, id string) (*SubModules, error) 
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]SubModules, error) 
	Update(ctx context.Context, ug *SubModules) error 
}

type subModuleRepository struct {
	DB *sql.DB
	log *zap.Logger
}

func NewSubModuleRepository(db *sql.DB, log *zap.Logger) SubModuleRepository {
	return &subModuleRepository{db,log}
}

// Create
func (r *subModuleRepository) Create(ctx context.Context, ug *SubModules) error {
	query := "INSERT INTO sub_modules (name) VALUES (?)"
	_, err := r.DB.ExecContext(ctx, query, ug.Name)
	if err != nil {
		return err
	}
	return nil
}

// Read (Get by ID)
func (r *subModuleRepository) GetByID(ctx context.Context, id string) (*SubModules, error) {
	query := "SELECT id, name FROM sub_modules WHERE id = ?"
	row := r.DB.QueryRowContext(ctx, query, id)
	var ug SubModules
	if err := row.Scan(&ug.ID, &ug.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ug, nil
}

// Read (Get All)
func (r *subModuleRepository) GetAll(ctx context.Context) ([]SubModules, error) {
	query := "SELECT id, name FROM sub_modules"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []SubModules
	for rows.Next() {
		var ug SubModules
		if err := rows.Scan(&ug.ID, &ug.Name); err != nil {
			return nil, err
		}
		groups = append(groups, ug)
	}
	return groups, nil
}

// Update
func (r *subModuleRepository) Update(ctx context.Context, ug *SubModules) error {
	query := "UPDATE sub_modules SET name = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, ug.Name, ug.ID)
	return err
}

// Delete
func (r *subModuleRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM sub_modules WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}