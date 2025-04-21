// user.go

package domain

type User struct {
    ID       int
    Username string
    Password string
}

type UserRepository interface {
    GetByUsername(username string) (*User, error)
}
