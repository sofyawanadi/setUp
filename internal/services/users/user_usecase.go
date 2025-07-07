package services

import (
	"setUp/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserUsecaseInterface interface {
	GetByEmail(c *gin.Context,email string) (*User, error)
	Login(c *gin.Context,email, password string) (*User, error) 
	InsertLogLogin(c *gin.Context, email string, success bool) error
	GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User,int64, error)
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

func (r *userUsecase) GetByEmail(c *gin.Context,email string) (*User, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userUsecase) Login(c *gin.Context,email, password string) (*User, error) {
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

func (r *userUsecase) GetAllUsers(c *gin.Context, params utils.QueryParams) ([]User,int64, error) {
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
	return users,count, nil
}