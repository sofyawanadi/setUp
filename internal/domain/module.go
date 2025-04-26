// user.go

package domain

type Module struct {
	BaseModel
	Name string `gorm:"null"`
	Const string `gorm:"not null"`
	SubModule []SubModule
}

type ModuleRepository interface {
	GetModule(username string) (*User, error)
}
