// user_usecase.go
package repository

import "project-name/internal/domain"

type UserMemoryRepo struct {
    users map[string]domain.User
}

func NewUserMemoryRepo() *UserMemoryRepo {
    return &UserMemoryRepo{
        users: map[string]domain.User{
            "admin": {ID: 1, Username: "admin", Password: "admin"},
        },
    }
}

func (r *UserMemoryRepo) GetByUsername(username string) (*domain.User, error) {
    user, ok := r.users[username]
    if !ok {
        return nil, nil
    }
    return &user, nil
}
