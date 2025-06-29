// user_usecase.go
package usecase

import "setUp/internal/domain"
import "setUp/internal/repository"

type UserUsecaseInterface interface {
	GetByUsername(username string) (*domain.User, error)
	Login(username, password string) (*domain.User, error)
}

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (r *UserUsecase) GetByUsername(username string) (*domain.User, error) {
	user, err := r.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserUsecase) Login(username, password string) (*domain.User, error) {
	user, err := r.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Password != password {
		return nil, nil
	}
	return user, nil
}
