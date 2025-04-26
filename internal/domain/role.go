// user.go

package domain

type Role struct {
	BaseModel
	Name string `gorm:"null"`
	Const string `gorm:"not null"`
	Permissions []Permission `gorm:"foreignKey:RoleId"`
}

type RoleRepository interface {
	GetRole(username string) (*Role, error)
}
