package services

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ModuleUsecase interface {
	GetByID(c *gin.Context, id int) (*Module, error)
	Create(c *gin.Context, module *Module) error
	Update(c *gin.Context, module *Module) error
	Delete(c *gin.Context, id int) error
	GetAll(c *gin.Context) ([]Module, error)
}

type moduleUsecase struct {
	repo ModuleRepository
	log  *zap.Logger
}

func NewModuleUsecase(repo ModuleRepository, log *zap.Logger) ModuleUsecase {
	return &moduleUsecase{repo: repo, log: log}
}

func (u *moduleUsecase) GetByID(c *gin.Context, id int) (*Module, error) {
	repoModule, err := u.repo.GetByID(c, int64(id))
	if err != nil {
		return nil, err
	}
	if repoModule == nil {
		return nil, nil
	}
	return repoModule, nil
}

func (u *moduleUsecase) Create(c *gin.Context, module *Module) error {
	return u.repo.Create(c, module)
}

func (u *moduleUsecase) Update(c *gin.Context, module *Module) error {
	return u.repo.Update(c, module)
}

func (u *moduleUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *moduleUsecase) GetAll(c *gin.Context) ([]Module, error) {
	repoModules, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return repoModules, nil
}
