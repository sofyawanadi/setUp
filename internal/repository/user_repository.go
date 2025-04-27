// user_repository.go
package repository

import (
    "setUp/internal/domain"
    "gorm.io/gorm"
)

type UserPostgres struct {
    DB *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
    return &UserPostgres{DB: db}
}

func (r *UserPostgres) GetByUsername(username string) (*domain.User, error) {
    var u domain.User
    result := r.DB.Where("username = ?", username).First(&u)
    if result.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &u, result.Error
}
