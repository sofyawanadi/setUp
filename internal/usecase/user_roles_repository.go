package usecase

import "setUp/internal/repository"
import "go.uber.org/zap"
import "github.com/gin-gonic/gin"

type UserRoles struct {
	ID   int
	Name string
}

type UserRolesUsecase interface {
	GetByID(c *gin.Context,id int) (*UserRoles, error)
	Create(c *gin.Context,userRoles *UserRoles) error
	Update(c *gin.Context,userRoles *UserRoles) error
	Delete(c *gin.Context,id int) error
	GetAll(c *gin.Context,) ([]*UserRoles, error)
}

type userRolesUsecase struct {
	repo repository.SubModuleRepository
	log      *zap.Logger
}


func NewUserRolesUsecase(repo repository.SubModuleRepository,log *zap.Logger) UserRolesUsecase {
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
	return &UserRoles{
		ID:   int(repoUserRoles.ID),
		Name: repoUserRoles.Name,
	}, nil
}

func (u *userRolesUsecase) Create(c *gin.Context, userRoles *UserRoles) error {
	repoUserRoles := &repository.SubModule{
		ID:   int64(userRoles.ID),
		Name: userRoles.Name,
	}
	return u.repo.Create(c, repoUserRoles)
}

func (u *userRolesUsecase) Update(c *gin.Context, userRoles *UserRoles) error {
	repoUserRoles := &repository.SubModule{
		ID:   int64(userRoles.ID),
		Name: userRoles.Name,
	}
	return u.repo.Update(c, repoUserRoles)
}

func (u *userRolesUsecase) Delete(c *gin.Context, id int) error {
	return u.repo.Delete(c, int64(id))
}

func (u *userRolesUsecase) GetAll(c *gin.Context) ([]*UserRoles, error) {
	repoUserRoless, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	userRoless := make([]*UserRoles, len(repoUserRoless))
	for i, repoUserRoles := range repoUserRoless {
		userRoless[i] = &UserRoles{
			ID:   int(repoUserRoles.ID),
			Name: repoUserRoles.Name,
		}
	}
	return userRoless, nil
}