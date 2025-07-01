// user_repository.go
package repository

import (
    "setUp/internal/domain"
    "gorm.io/gorm"
    "go.uber.org/zap"
    "github.com/gin-gonic/gin"
)

type UserRepository interface {
    GetByEmail(email string) (*domain.User, error)
    InsertLogLogins(c *gin.Context, email string, success bool) error
}

type userRepository struct {
	db *gorm.DB
    log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{db,log}
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var u domain.User
	result := r.db.Where("email = ?", email).First(&u)
	if result.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &u, result.Error
}

func (r *userRepository) InsertLogLogins(c *gin.Context,Email string, success bool) error {
    entry := domain.LogLoginModel{
        Email:  Email,
        ClientIP: c.ClientIP(),
    }
    if err := r.db.Create(&entry).Error; err != nil {
        r.log.Error("failed to log login attempt", zap.Error(err))
        return err
    }
    return nil
}