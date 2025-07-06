package services

import ( 
	"setUp/internal/domain"
)

type UserRole struct {
	domain.BaseModel
	UserId string `gorm:"column:user_id;not null"`
	RoleId string `gorm:"column:role_id;not null"`
}

func (UserRole) TableName() string {
    return "user_roles"
}
