package usecase

import "setUp/internal/repository"
import "go.uber.org/zap"
import "github.com/gin-gonic/gin"

type Module struct {
	ID   int
	Name string
}

type ModuleUsecase interface {
	GetByID(c *gin.Context,id int) (*Module, error)
	Create(c *gin.Context,module *Module) error
	Update(c *gin.Context,module *Module) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]*Module, error)
}

type moduleUsecase struct {
	repo repository.ModuleRepository
	log      *zap.Logger
}


func NewModuleUsecase(repo repository.ModuleRepository,log *zap.Logger) ModuleUsecase {
	return &moduleUsecase{repo: repo, log:log,}
}

func (u *moduleUsecase) GetByID(c *gin.Context, id int) (*Module, error) {
	repoModule, err := u.repo.GetByID(c, int64(id))
	if err != nil {
		return nil, err
	}
	if repoModule == nil {
		return nil, nil
	}
	return &Module{
		ID:   int(repoModule.ID),
		Name: repoModule.Name,
	}, nil
}

func (u *moduleUsecase) Create(c *gin.Context, module *Module) error {
	repoModule := &repository.Module{
		ID:   int64(module.ID),
		Name: module.Name,
	}
	return u.repo.Create(c, repoModule)
}

func (u *moduleUsecase) Update(c *gin.Context, module *Module) error {
	repoModule := &repository.Module{
		ID:   int64(module.ID),
		Name: module.Name,
	}
	return u.repo.Update(c, repoModule)
}

func (u *moduleUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *moduleUsecase) GetAll(c *gin.Context) ([]*Module, error) {
	repoModules, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	modules := make([]*Module, len(repoModules))
	for i, repoModule := range repoModules {
		modules[i] = &Module{
			ID:   int(repoModule.ID),
			Name: repoModule.Name,
		}
	}
	return modules, nil
}