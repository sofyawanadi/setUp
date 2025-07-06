package services
import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

type UserRoles struct {
	ID   int64
	Name string
}

type UserRolesRepository interface {
	Create(ctx context.Context, ug *UserRoles) error
	GetByID(ctx context.Context, id int64) (*UserRoles, error) 
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]UserRoles, error) 
	Update(ctx context.Context, ug *UserRoles) error 
}

type userRolesRepository struct {
	DB *sql.DB
	log *zap.Logger
}

func NewUserRolesRepository(db *sql.DB, log *zap.Logger) UserRolesRepository {
	return &userRolesRepository{db,log}
}

// Create
func (r *userRolesRepository) Create(ctx context.Context, ug *UserRoles) error {
	query := "INSERT INTO user_roles (name) VALUES (?)"
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
func (r *userRolesRepository) GetByID(ctx context.Context, id int64) (*UserRoles, error) {
	query := "SELECT id, name FROM user_roles WHERE id = ?"
	row := r.DB.QueryRowContext(ctx, query, id)
	var ug UserRoles
	if err := row.Scan(&ug.ID, &ug.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ug, nil
}

// Read (Get All)
func (r *userRolesRepository) GetAll(ctx context.Context) ([]UserRoles, error) {
	query := "SELECT id, name FROM user_roles"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []UserRoles
	for rows.Next() {
		var ug UserRoles
		if err := rows.Scan(&ug.ID, &ug.Name); err != nil {
			return nil, err
		}
		groups = append(groups, ug)
	}
	return groups, nil
}

// Update
func (r *userRolesRepository) Update(ctx context.Context, ug *UserRoles) error {
	query := "UPDATE user_roles SET name = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, ug.Name, ug.ID)
	return err
}

// Delete
func (r *userRolesRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM user_roles WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}