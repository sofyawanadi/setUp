// user.go

package domain

type UserRole struct {
	BaseModel
	UserId string `gorm:"column:user_id;not null"`
	RoleId string `gorm:"column:role_id;not null"`
}

func (UserRole) TableName() string {
    return "user_roles"
}

type UserRoleRepository interface {
	GetAll() (*UserRole, error)
}
