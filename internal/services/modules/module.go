// user.go
package services

import (
	"setUp/internal/domain"
	sub "setUp/internal/services/sub_modules"
)

type Module struct {
	domain.BaseModel
	Name        string `gorm:"null;column:name"`
	Const       string `gorm:"not null;column:const"`
	Description string `gorm:"null;column:description"`
	SubModule   []sub.SubModules
}

func (Module) TableName() string {
	return "modules"
}
