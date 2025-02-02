package usecase

import (
	"github.com/stretchr/testify/assert"
	"personal_security/internal/domain"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	controller := NewUser()

	newUser := &domain.User{
		Name:     "Иван Иванов",
		Email:    "ivan@example.com",
		Password: "securepassword123",
	}

	userID := controller.RegisterUser(newUser)

	assert.Equal(t, userID, 1)

	assert.Equal(t, len(users), 1)
}

func TestLoginUser(t *testing.T) {
	controller := NewUser()

	testUser := &domain.User{
		Name:     "Иван Иванов",
		Email:    "ivan@example.com",
		Password: "securepassword123",
	}
	controller.RegisterUser(testUser)

	loginUser := &domain.User{
		Email:    "ivan@example.com",
		Password: "securepassword123",
	}

	loggedInUser, err := controller.LoginUser(loginUser)

	assert.Nil(t, err)
	assert.Equal(t, loggedInUser.Email, loginUser.Email)

	invalidUser := &domain.User{
		Email:    "ivan@example.com",
		Password: "wrongpassword",
	}
	_, err = controller.LoginUser(invalidUser)

	assert.NotNil(t, err)
}
