// user_repository.go
package repository

import (
    "setUp/internal/domain"
    "gorm.io/gorm"
)

type UserRepository interface {
    GetByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var u domain.User
	result := r.db.Where("username = ?", username).First(&u)
	if result.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &u, result.Error
}
