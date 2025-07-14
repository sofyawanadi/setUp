package services

import (
	"setUp/internal/domain"
	"setUp/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
    GetByEmail(email string) (*User, error)
    GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User, error)
    InsertLogLogins(c *gin.Context, email string, success bool) error
    GetCountUsers(c *gin.Context, params utils.QueryParams) (int64, error)
    GetByID(id string) (*User, error)
    GetByUsername(username string) (*User, error) 
    InsertUser(c *gin.Context, user PostUserRequest) (*User, error)
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

func (r *userRepository) GetByUsername(username string) (*User, error) {
	var u User
	result := r.db.Where("username = ?", username).First(&u)
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

func (r *userRepository) GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User, error) {
    var users []User
    result := r.db.Select("id", "email", "username", "created_at", "updated_at")
    query := utils.ApplyQuery(result, params)
    if err := query.Find(&users).Error; err != nil {
        r.log.Error("failed to get all users", zap.Error(err))
        return nil, err
    }

    return users, nil
}

func (r *userRepository) GetCountUsers(c *gin.Context, params utils.QueryParams) (int64, error) {
    var count int64
    result := r.db.Model(&User{})
    query := utils.ApplyQuery(result, params)
    if err := query.Count(&count).Error; err != nil {
        r.log.Error("failed to get user count", zap.Error(err))
        return 0, err
    }
    return count, nil
}
func (r *userRepository) GetByID(id string) (*User, error) {
    var u User
    result := r.db.Where("id = ?", id).First(&u)
    if result.Error == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &u, result.Error
}

func (r *userRepository) InsertUser(c *gin.Context, user PostUserRequest) (*User, error) {
    usrId := c.GetString("userID")
    usr := User{
        Username: user.Username,
        Password: user.Password,
        Email:    user.Email,
    }
    usr.CreatedBy = usrId
    result := r.db.Create(&usr)
    if result.Error != nil {
        r.log.Error("failed to get user count", zap.Error(result.Error))
	} else {
        r.log.Info("inserted id user", zap.String("user_id", usr.ID.String()))
	}
    return &usr, result.Error
}