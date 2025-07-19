package services

import (
	"database/sql"
	"fmt"
	"setUp/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NoteUsecase interface {
	GetByID(c *gin.Context, id string) (*Note, error)
	Create(c *gin.Context, note *NoteRequest) (*Note, error)
	Update(c *gin.Context, note *NoteUpdateRequest) (*Note, error)
	Delete(c *gin.Context, id string) error
	GetAll(c *gin.Context, param utils.QueryParams) ([]Note, int64, error)
}

type noteUsecase struct {
	repo NoteRepository
	log  *zap.Logger
}

func NewNoteUsecase(repo NoteRepository, log *zap.Logger) NoteUsecase {
	return &noteUsecase{repo: repo, log: log}
}

func (u *noteUsecase) GetByID(c *gin.Context, id string) (*Note, error) {
	repoNote, err := u.repo.GetByID(c, id)
	if err != nil {
		return nil, fmt.Errorf("note not found")
	}
	if repoNote == nil {
		return nil, fmt.Errorf("note not found")
	}
	return repoNote, nil
}

func (u *noteUsecase) Create(c *gin.Context, note *NoteRequest) (*Note, error) {
	return u.repo.Create(c, note)
}

func (u *noteUsecase) Update(c *gin.Context, note *NoteUpdateRequest) (*Note, error) {
	return u.repo.Update(c, note)
}

func (u *noteUsecase) Delete(c *gin.Context, id string) error {
	repoNote, err := u.repo.GetByID(c, id)
	if err != nil {
		return err
	}
	if repoNote == nil {
		return fmt.Errorf("notes not found")
	}
	usrId := c.GetString("userID")
	repoNote.DeletedBy = usrId
	repoNote.DeletedAt = sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}
	repoNote.IsActive = false
	err = u.repo.Delete(c, *repoNote)
	if err != nil {
		return err
	}

	return err
}

func (u *noteUsecase) GetAll(c *gin.Context, param utils.QueryParams) ([]Note, int64, error) {
	repoNotes, err := u.repo.GetAll(c, param)
	if err != nil {
		return nil, 0, err
	}
	count, err := u.repo.GetCount(c, param)
	if err != nil {
		return nil, 0, err
	}
	return repoNotes, count, nil
}
