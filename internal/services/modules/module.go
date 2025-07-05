// user.go

package domain

type Module struct {
	BaseModel
	Name string `gorm:"null;column:name"`
	Const string `gorm:"not null;column:const"`
	Description string `gorm:"null;column:description"`
	SubModule []SubModule
}

func (Module) TableName() string {
	return "modules"
}

type ModuleRepository interface {
	GetAll() ([]*Module, error)
}
