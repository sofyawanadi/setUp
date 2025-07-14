package domain

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey;column:id"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	DeletedAt sql.NullTime `gorm:"index;column:deleted_at"`
	CreatedBy string       `gorm:"column:created_by;null"`
	UpdatedBy string       `gorm:"column:updated_by;null"`
	DeletedBy string       `gorm:"column:deleted_by;null"`
	IsActive  bool         `gorm:"column:is_active;null"`
	// gorm.Model
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
