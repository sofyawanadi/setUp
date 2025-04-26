// user.go

package domain

type UserRole struct {
	BaseModel
	UserId string `gorm:"not null"`
	RoleId string `gorm:"not null"`
}

type UserRoleRepository interface {
	GetBy(username string) (*UserRole, error)
}
