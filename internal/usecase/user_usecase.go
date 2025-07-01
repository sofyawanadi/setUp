// user_usecase.go
package usecase

import "setUp/internal/domain"
import "setUp/internal/repository"
import "go.uber.org/zap"
import "github.com/gin-gonic/gin"
type UserUsecaseInterface interface {
	GetByUsername(c *gin.Context,username string) (*domain.User, error)
	Login(c *gin.Context,username, password string) (*domain.User, error)
	InsertLogLogin(c *gin.Context, username string, success bool)
}

type UserUsecase struct {
	userRepo repository.UserRepository
	log      *zap.Logger
}

func NewUserUsecase(userRepo repository.UserRepository, log *zap.Logger) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		log:      log,
	}
}

func (r *UserUsecase) GetByEmail(c *gin.Context,email string) (*domain.User, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserUsecase) Login(c *gin.Context,email, password string) (*domain.User, error) {
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