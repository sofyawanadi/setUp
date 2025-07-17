package services

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SubModulesUsecase interface {
	GetByID(c *gin.Context, id string) (*SubModules, error)
	Create(c *gin.Context, subModule SubModules) error
	Update(c *gin.Context, subModule SubModules) error
	Delete(c *gin.Context, id string) error
	GetAll(c *gin.Context) ([]SubModules, error)
}

type subModulesUsecase struct {
	repo SubModuleRepository
	log  *zap.Logger
}

func NewSubModulesUsecase(repo SubModuleRepository, log *zap.Logger) SubModulesUsecase {
	return &subModulesUsecase{repo: repo, log: log}
}

func (u *subModulesUsecase) GetByID(c *gin.Context, id string) (*SubModules, error) {
	repoSubModule, err := u.repo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	if repoSubModule == nil {
		return nil, nil
	}
	return repoSubModule, nil
}
func (u *subModulesUsecase) Create(c *gin.Context, subModule SubModules) error {
	return u.repo.Create(c, &subModule)
}

func (u *subModulesUsecase) Update(c *gin.Context, subModule SubModules) error {
	return u.repo.Update(c, &subModule)
}

func (u *subModulesUsecase) Delete(c *gin.Context, id string) error {
	return u.repo.Delete(c, id)
}

func (u *subModulesUsecase) GetAll(c *gin.Context) ([]SubModules, error) {
	repoSubModulesList, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	if repoSubModulesList == nil {
		return nil, nil
	}
	return repoSubModulesList, nil
}
