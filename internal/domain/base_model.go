package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey;column:id"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	DeletedAt sql.NullTime `gorm:"index;column:deleted_at"`
	CreatedBy string       `gorm:"column:created_by;null"`
	UpdatedBy string       `gorm:"column:updated_by;null"`
	DeletedBy string       `gorm:"column:deleted_by;null"`
	IsActive  bool         `gorm:"column:is_active;default:true"`
	// gorm.Model
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}

type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}