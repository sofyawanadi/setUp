// user_repository.go
package repository

import (
    "setUp/internal/domain"
    "gorm.io/gorm"
    "go.uber.org/zap"
)

type UserRepository interface {
    GetByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
    log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{db,log}
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var u domain.User
	result := r.db.Where("username = ?", username).First(&u)
	if result.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &u, result.Error
}
