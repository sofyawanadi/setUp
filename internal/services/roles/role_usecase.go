package services

import "go.uber.org/zap"
import "github.com/gin-gonic/gin"


type RoleUsecase interface {
	GetByID(c *gin.Context,id int) (*Role, error)
	Create(c *gin.Context,role *Role) error
	Update(c *gin.Context,role *Role) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]Role, error)
}

type roleUsecase struct {
	repo RoleRepository
	log      *zap.Logger
}


func NewRoleUsecase(repo RoleRepository,log *zap.Logger) RoleUsecase {
	return &roleUsecase{repo: repo, log:log,}
}

func (u *roleUsecase) GetByID(c *gin.Context, id int) (*Role, error) {
	repoRole, err := u.repo.GetByID(c, int64(id))
	if err != nil {
		return nil, err
	}
	if repoRole == nil {
		return nil, nil
	}
	return repoRole, nil
}

func (u *roleUsecase) Create(c *gin.Context, role *Role) error {
	return u.repo.Create(c, role)
}

func (u *roleUsecase) Update(c *gin.Context, role *Role) error {
	return u.repo.Update(c, role)
}

func (u *roleUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *roleUsecase) GetAll(c *gin.Context) ([]Role, error) {
	repoRoles, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}

	return repoRoles, nil
}