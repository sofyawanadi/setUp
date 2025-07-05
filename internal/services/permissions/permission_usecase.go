package usecase

import "setUp/internal/repository"
import "go.uber.org/zap"
import "github.com/gin-gonic/gin"

type Permissions struct {
	ID   int
	Name string
}

type PermissionsUsecase interface {
	GetByID(c *gin.Context,id int) (*Permissions, error)
	Create(c *gin.Context,permissions *Permissions) error
	Update(c *gin.Context,permissions *Permissions) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]*Permissions, error)
}

type permissionsUsecase struct {
	repo repository.PermissionsRepository
	log      *zap.Logger
}


func NewPermissionsUsecase(repo repository.PermissionsRepository,log *zap.Logger) PermissionsUsecase {
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
	return &Permissions{
		ID:   int(repoPermissions.ID),
		Name: repoPermissions.Name,
	}, nil
}

func (u *permissionsUsecase) Create(c *gin.Context, permissions *Permissions) error {
	repoPermissions := &repository.Permissions{
		ID:   int64(permissions.ID),
		Name: permissions.Name,
	}
	return u.repo.Create(c, repoPermissions)
}

func (u *permissionsUsecase) Update(c *gin.Context, permissions *Permissions) error {
	repoPermissions := &repository.Permissions{
		ID:   int64(permissions.ID),
		Name: permissions.Name,
	}
	return u.repo.Update(c, repoPermissions)
}

func (u *permissionsUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *permissionsUsecase) GetAll(c *gin.Context) ([]*Permissions, error) {
	repoPermissionss, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	permissionss := make([]*Permissions, len(repoPermissionss))
	for i, repoPermissions := range repoPermissionss {
		permissionss[i] = &Permissions{
			ID:   int(repoPermissions.ID),
			Name: repoPermissions.Name,
		}
	}
	return permissionss, nil
}