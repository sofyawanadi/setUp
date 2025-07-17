package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogLoginModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;column:id"`
	Email     string    `gorm:"column:email"`
	ClientIP  string    `gorm:"column:client_ip"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (LogLoginModel) TableName() string {
	return "log_logins"
}

func (base *LogLoginModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
