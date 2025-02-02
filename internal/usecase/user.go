package usecase

import (
	"errors"
	"personal_security/internal/domain"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

var users []*domain.User

func (u *User) RegisterUser(newUser *domain.User) int {
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	return newUser.ID
}

func (u *User) LoginUser(loginUser *domain.User) (*domain.User, error) {
	for _, user := range users {
		if user.Email == loginUser.Email && user.Password == loginUser.Password {
			return user, nil
		}
	}
	return nil, errors.New("no user found")
}
