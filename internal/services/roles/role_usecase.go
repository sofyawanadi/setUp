package usecase

import "setUp/internal/repository"
import "go.uber.org/zap"
import "github.com/gin-gonic/gin"

type Role struct {
	ID   int
	Name string
}

type RoleUsecase interface {
	GetByID(c *gin.Context,id int) (*Role, error)
	Create(c *gin.Context,role *Role) error
	Update(c *gin.Context,role *Role) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]*Role, error)
}

type roleUsecase struct {
	repo repository.RoleRepository
	log      *zap.Logger
}


func NewRoleUsecase(repo repository.RoleRepository,log *zap.Logger) RoleUsecase {
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
	return &Role{
		ID:   int(repoRole.ID),
		Name: repoRole.Name,
	}, nil
}

func (u *roleUsecase) Create(c *gin.Context, role *Role) error {
	repoRole := &repository.Role{
		ID:   int64(role.ID),
		Name: role.Name,
	}
	return u.repo.Create(c, repoRole)
}

func (u *roleUsecase) Update(c *gin.Context, role *Role) error {
	repoRole := &repository.Role{
		ID:   int64(role.ID),
		Name: role.Name,
	}
	return u.repo.Update(c, repoRole)
}

func (u *roleUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *roleUsecase) GetAll(c *gin.Context) ([]*Role, error) {
	repoRoles, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	roles := make([]*Role, len(repoRoles))
	for i, repoRole := range repoRoles {
		roles[i] = &Role{
			ID:   int(repoRole.ID),
			Name: repoRole.Name,
		}
	}
	return roles, nil
}