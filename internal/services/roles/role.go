package services

import (
	"setUp/internal/domain"
	per "setUp/internal/services/permissions"
)

type Role struct {
	domain.BaseModel
	Name        string            `gorm:"column:name;null"`
	Const       string            `gorm:"column:const;not null"`
	Permissions []per.Permissions `gorm:"foreignKey:RoleId"`
}

func (Role) TableName() string {
	return "roles"
}
