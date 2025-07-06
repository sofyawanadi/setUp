package services

import "go.uber.org/zap"
import "github.com/gin-gonic/gin"


type UserRolesUsecase interface {
	GetByID(c *gin.Context,id int) (*UserRoles, error)
	Create(c *gin.Context,userRoles *UserRoles) error
	Update(c *gin.Context,userRoles *UserRoles) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]UserRoles, error)
}

type userRolesUsecase struct {
	repo UserRolesRepository
	log      *zap.Logger
}


func NewUserRolesUsecase(repo UserRolesRepository,log *zap.Logger) UserRolesUsecase {
	return &userRolesUsecase{repo: repo, log:log,}
}

func (u *userRolesUsecase) GetByID(c *gin.Context, id int) (*UserRoles, error) {
	repoUserRoles, err := u.repo.GetByID(c, int64(id))
	if err != nil {
		return nil, err
	}
	if repoUserRoles == nil {
		return nil, nil
	}
	return repoUserRoles, nil
}

func (u *userRolesUsecase) Create(c *gin.Context, userRoles *UserRoles) error {

	return u.repo.Create(c, userRoles)
}

func (u *userRolesUsecase) Update(c *gin.Context, userRoles *UserRoles) error {
	return u.repo.Update(c, userRoles)
}

func (u *userRolesUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *userRolesUsecase) GetAll(c *gin.Context) ([]UserRoles, error) {
	repoUserRoles, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}

	return repoUserRoles, nil
}