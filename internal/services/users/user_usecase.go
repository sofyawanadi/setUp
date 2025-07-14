package services

import (
	"fmt"
	"os"
	"setUp/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserUsecaseInterface interface {
	GetByEmail(c *gin.Context, email string) (*User, error)
	Login(c *gin.Context, email, password string) (*User, error)
	InsertLogLogin(c *gin.Context, email string, success bool) error
	GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User, int64, error)
	GetByID(c *gin.Context) (*User, error)
	InsertUser(c *gin.Context, user PostUserRequest) (*User, error)
}

type userUsecase struct {
	userRepo UserRepository
	log      *zap.Logger
}

func NewUserUsecase(userRepo UserRepository, log *zap.Logger) UserUsecaseInterface {
	return &userUsecase{
		userRepo: userRepo,
		log:      log,
	}
}

func (r *userUsecase) GetByEmail(c *gin.Context, email string) (*User, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userUsecase) Login(c *gin.Context, email, password string) (*User, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Password != password {
		return nil, nil
	}
	return user, nil
}

func (r *userUsecase) InsertLogLogin(c *gin.Context, email string, success bool) error {
	err := r.userRepo.InsertLogLogins(c, email, success)
	if err != nil {
		r.log.Error("failed to log login attempt", zap.Error(err))
		return err
	}
	return nil
}

func (r *userUsecase) GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User, int64, error) {
	users, err := r.userRepo.GetAllUsers(c, params)
	if err != nil {
		r.log.Error("failed to get all users", zap.Error(err))
		return nil, 0, err
	}
	count, err := r.userRepo.GetCountUsers(c, params)
	if err != nil {
		r.log.Error("failed to get user count", zap.Error(err))
		return nil, 0, err
	}
	return users, count, nil
}

func (r *userUsecase) GetByID(c *gin.Context) (*User, error) {
	userId, exists := c.Get("userID")
	if !exists {
		r.log.Error("failed to get user ID from context")
		return nil, fmt.Errorf("failed to get user ID from context")
	}
	userIDStr, ok := userId.(string)
	if !ok {
		r.log.Error("user ID in context is not a string")
		return nil, fmt.Errorf("user ID in context is not a string")
	}
	user, err2 := r.userRepo.GetByID(userIDStr)
	if err2 != nil {
		return nil, err2
	}
	return user, nil
}

func (r *userUsecase) InsertUser(c *gin.Context, userPost PostUserRequest) (*User, error) {

	checkEmail, err := r.userRepo.GetByEmail(userPost.Email)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	checkUsername, err := r.userRepo.GetByUsername(userPost.Username)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	if checkEmail != nil && checkUsername != nil {
		return nil, fmt.Errorf("username or email is used")
	}
	hashPassword, err := utils.HashPassword(userPost.Password)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	userPost.Password = hashPassword
	usr, err := r.userRepo.InsertUser(c, userPost)
	if err != nil {
		return nil, err
	}
	appUrl := os.Getenv("APP_URL")
	year :=time.Now().Year()
	utils.SendMail([]string{usr.Email}, "Notification New User", "create_user.html", map[string]interface{}{
		"Username": usr.Username,
		"Year":     year,
		"LoginURL": appUrl,
	})
	return usr, nil
}
