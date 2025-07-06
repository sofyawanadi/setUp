package services

import "go.uber.org/zap"
import "github.com/gin-gonic/gin"

type PermissionsUsecase interface {
	GetByID(c *gin.Context,id int) (*Permissions, error)
	Create(c *gin.Context,permissions *Permissions) error
	Update(c *gin.Context,permissions *Permissions) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]Permissions, error)
}

type permissionsUsecase struct {
	repo PermissionsRepository
	log      *zap.Logger
}


func NewPermissionsUsecase(repo PermissionsRepository,log *zap.Logger) PermissionsUsecase {
	return &permissionsUsecase{repo: repo, log:log,}
}

func (u *permissionsUsecase) GetByID(c *gin.Context, id int) (*Permissions, error) {
	repoPermissions, err := u.repo.GetByID(c, int64(id))
	if err != nil {
		return nil, err
	}
	if repoPermissions == nil {
		return nil, nil
	}
	return repoPermissions, nil
}

func (u *permissionsUsecase) Create(c *gin.Context, permissions *Permissions) error {
	return u.repo.Create(c, permissions)
}

func (u *permissionsUsecase) Update(c *gin.Context, permissions *Permissions) error {
	return u.repo.Update(c, permissions)
}

func (u *permissionsUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *permissionsUsecase) GetAll(c *gin.Context) ([]Permissions, error) {
	repoPermissions, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return repoPermissions, nil
}