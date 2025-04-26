// user.go

package domain

type User struct {
	BaseModel
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Address  string `gorm:"null"`
	Token    string `gorm:"-"`
}

type UserRepository interface {
	GetByUsername(username string) (*User, error)
}
