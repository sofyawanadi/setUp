package services

import (
	// "setUp/internal/domain"
	// "setUp/internal/utils"

	// "github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UploadRepository interface {
}

type uploadRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUploadRepository(db *gorm.DB, log *zap.Logger) UploadRepository {
	return &uploadRepository{db, log}
}
