// user_repository.go
package repository

import (
    "database/sql"
    "project-name/internal/domain"

    "github.com/jmoiron/sqlx"
)

type UserPostgres struct {
    DB *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
    return &UserPostgres{DB: db}
}

func (r *UserPostgres) GetByUsername(username string) (*domain.User, error) {
    var u domain.User
    err := r.DB.Get(&u, "SELECT id, username, password FROM users WHERE username=$1", username)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &u, err
}
