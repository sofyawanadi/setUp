package services
import (
    "setUp/internal/domain"
    "gorm.io/gorm"
    "go.uber.org/zap"
    "github.com/gin-gonic/gin"
)

type UserRepository interface {
    GetByEmail(email string) (*User, error)
    GetAllUsers() ([]User, error) 
    InsertLogLogins(c *gin.Context, email string, success bool) error
}

type userRepository struct {
	db *gorm.DB
    log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{db,log}
}

func (r *userRepository) GetByEmail(email string) (*User, error) {
	var u User
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

func (r *userRepository) GetAllUsers() ([]User, error) {
    var users []User
    result := r.db.Select("id", "email", "username", "created_at", "updated_at").Find(&users)
    if result.Error != nil {
        r.log.Error("failed to get all users", zap.Error(result.Error))
        return nil, result.Error
    }
    return users, nil
}