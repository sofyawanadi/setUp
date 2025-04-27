// user.go

package domain

type Role struct {
	BaseModel
	Name        string       `gorm:"column:name;null"`
	Const       string       `gorm:"column:const;not null"`
	Permissions []Permission `gorm:"foreignKey:RoleId"`
}

func (Role) TableName() string {
    return "roles"
}

type RoleRepository interface {
	GetRole(username string) (*Role, error)
}
