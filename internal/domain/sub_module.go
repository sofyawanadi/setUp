// user.go

package domain

type SubModule struct {
	BaseModel
	Name string `gorm:"null"`
	Const string `gorm:"not null"`
	Permissions []Permission `gorm:"foreignKey:ModuleId"`
}

type SubModuleRepository interface {
	GetModule(module string) (*SubModule, error)
}
