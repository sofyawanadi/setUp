// user.go

package domain

import (
	// "golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"column:username;unique;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Address  string `gorm:"column:address;null"`
	Token    string `gorm:"-"`
}
func (User) TableName() string {
    return "users"
}

// func (u *User) BeforeSave(tx *gorm.DB) (err error) {
// 	// Cek apakah kolom Password di-update
// 	if tx.Statement.Changed("Password") {
// 		// Hash password baru
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			return err
// 		}
// 		u.Password = string(hashedPassword)
// 	}
// 	return nil
// }

type UserRepository interface {
	GetByUsername(username string) (*User, error)
}
