package domain

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	CreatedBy string
	UpdatedBy string
	DeletedBy string
	IsActive  bool
	// gorm.Model
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
    base.ID = uuid.New()
    return
}