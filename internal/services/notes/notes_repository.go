package services

import (
	"context"
	"setUp/internal/domain"
	"setUp/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(ctx *gin.Context, note *NoteRequest) (*Note, error)
	GetByID(ctx *gin.Context, id string) (*Note, error)
	Delete(ctx *gin.Context, note Note) error
	GetAll(ctx context.Context, param utils.QueryParams) ([]Note, error)
	Update(ctx *gin.Context, note *NoteUpdateRequest) (*Note, error)
	GetCount(ctx context.Context, params utils.QueryParams) (int64, error)
}

type noteRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewNoteRepository(db *gorm.DB, log *zap.Logger) NoteRepository {
	return &noteRepository{db, log}
}

func (r *noteRepository) Create(ctx *gin.Context, note *NoteRequest) (*Note, error) {
	usrId := ctx.GetString("userID")
	create := Note{
		Title:   note.Title,
		Content: note.Content,
		BaseModel: domain.BaseModel{
			CreatedBy: usrId,
			UpdatedBy: usrId,
			IsActive:  true,
		},
	}
	if err := r.db.WithContext(ctx).Create(&create).Error; err != nil {
		r.log.Error("failed to insert note", zap.Error(err))
		return nil, err
	}
	return &create, nil
}

func (r *noteRepository) GetByID(ctx *gin.Context, id string) (*Note, error) {
	var note Note
	if err := r.db.WithContext(ctx).First(&note, "id = ? and is_active = ?", id, true).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Error("failed to retrieve note", zap.Error(err))
			return nil, err
		}
		r.log.Error("failed to retrieve note", zap.Error(err))
		return nil, err
	}
	return &note, nil
}

func (r *noteRepository) GetAll(ctx context.Context, params utils.QueryParams) ([]Note, error) {
	var notes []Note
	query := utils.ApplyQuery(r.db, params)
	query.WithContext(ctx).Where("is_active = ?", true)

	if err := query.Find(&notes).Error; err != nil {
		r.log.Error("failed to get all note", zap.Error(err))
		return nil, err
	}
	return notes, nil
}
func (r *noteRepository) GetCount(ctx context.Context, params utils.QueryParams) (int64, error) {
	var notes int64
	query := utils.ApplyQuery(r.db, params)
	query.Model(&Note{})
	query.WithContext(ctx).Where("is_active = ?", true)

	if err := query.Count(&notes).Error; err != nil {
		r.log.Error("failed to count note", zap.Error(err))
		return 0, err
	}
	return notes, nil
}

func (r *noteRepository) Update(ctx *gin.Context, note *NoteUpdateRequest) (*Note, error) {
	usrId := ctx.GetString("userID")
	update := Note{
		Title:   note.Title,
		Content: note.Content,
		BaseModel: domain.BaseModel{
			CreatedBy: usrId,
			UpdatedBy: usrId,
		},
	}
	update.ID = note.Id
	if err := r.db.WithContext(ctx).Save(&update).Error; err != nil {
		r.log.Error("failed to Update note", zap.Error(err))
		return nil, err
	}
	return &update, nil
}

func (r *noteRepository) Delete(ctx *gin.Context, note Note) error {
	return r.db.WithContext(ctx).Save(&note).Error
}
