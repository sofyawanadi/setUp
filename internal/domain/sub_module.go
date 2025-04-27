// user.go

package domain

type SubModule struct {
	BaseModel
	ModuleId    string       `gorm:"column:module_id;not null"`
	Name        string       `gorm:"column:name;null"`
	Const       string       `gorm:"column:const;not null"`
	Description string       `gorm:"column:description;null"`
	Permissions []Permission `gorm:"foreignKey:ModuleId;column:permissions"`
}

func (SubModule) TableName() string {
    return "sub_modules"
}

type SubModuleRepository interface {
	GetSubModule(module string) (*SubModule, error)
}
