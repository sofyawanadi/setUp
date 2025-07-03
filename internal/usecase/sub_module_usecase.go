package usecase

import "setUp/internal/repository"
import "go.uber.org/zap"
import "github.com/gin-gonic/gin"

type SubModules struct {
	ID   int
	Name string
}

type SubModulesUsecase interface {
	GetByID(c *gin.Context,id int) (*SubModules, error)
	Create(c *gin.Context,subModules *SubModules) error
	Update(c *gin.Context,subModules *SubModules) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]*SubModules, error)
}

type subModulesUsecase struct {
	repo repository.SubModuleRepository
	log      *zap.Logger
}


func NewSubModulesUsecase(repo repository.SubModuleRepository,log *zap.Logger) SubModulesUsecase {
	return &subModulesUsecase{repo: repo, log:log,}
}

func (u *subModulesUsecase) GetByID(c *gin.Context, id int) (*SubModules, error) {
	repoSubModules, err := u.repo.GetByID(c, int64(id))
	if err != nil {
		return nil, err
	}
	if repoSubModules == nil {
		return nil, nil
	}
	return &SubModules{
		ID:   int(repoSubModules.ID),
		Name: repoSubModules.Name,
	}, nil
}

func (u *subModulesUsecase) Create(c *gin.Context, subModules *SubModules) error {
	repoSubModules := &repository.SubModule{
		ID:   int64(subModules.ID),
		Name: subModules.Name,
	}
	return u.repo.Create(c, repoSubModules)
}

func (u *subModulesUsecase) Update(c *gin.Context, subModules *SubModules) error {
	repoSubModules := &repository.SubModule{
		ID:   int64(subModules.ID),
		Name: subModules.Name,
	}
	return u.repo.Update(c, repoSubModules)
}

func (u *subModulesUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *subModulesUsecase) GetAll(c *gin.Context) ([]*SubModules, error) {
	repoSubModuless, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	subModuless := make([]*SubModules, len(repoSubModuless))
	for i, repoSubModules := range repoSubModuless {
		subModuless[i] = &SubModules{
			ID:   int(repoSubModules.ID),
			Name: repoSubModules.Name,
		}
	}
	return subModuless, nil
}