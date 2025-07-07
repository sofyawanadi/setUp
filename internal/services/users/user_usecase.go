package services

import (
	"setUp/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserUsecaseInterface interface {
	GetByUsername(c *gin.Context,username string) (*User, error)
	Login(c *gin.Context,username, password string) (*User, error)
	InsertLogLogin(c *gin.Context, username string, success bool)
}

type UserUsecase struct {
	userRepo UserRepository
	log      *zap.Logger
}

func NewUserUsecase(userRepo UserRepository, log *zap.Logger) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		log:      log,
	}
}

func (r *UserUsecase) GetByEmail(c *gin.Context,email string) (*User, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserUsecase) Login(c *gin.Context,email, password string) (*User, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Password != password {
		return nil, nil
	}
	return user, nil
}

func (r *UserUsecase) InsertLogLogin(c *gin.Context, username string, success bool) error {
	err := r.userRepo.InsertLogLogins(c, username, success)
	if err != nil {
		r.log.Error("failed to log login attempt", zap.Error(err))
		return err
	}
	return nil
}

func (r *UserUsecase) GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User, error) {
	users, err := r.userRepo.GetAllUsers(c, params)
	if err != nil {
		r.log.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	return users, nil
}