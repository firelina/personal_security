package usecase

import (
	"context"
	"errors"
	"personal_security/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) RegisterUser(ctx context.Context, newUser *domain.User) (int, error) {
	args := m.Called(ctx, newUser)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) LoginUser(ctx context.Context, loginUser *domain.User) (*domain.User, error) {
	args := m.Called(ctx, loginUser)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)
	t.Run("success register user", func(t *testing.T) {
		newUser := &domain.User{Name: "John Doe", Email: "john@example.com", Password: "password123"}

		mockRepo.On("RegisterUser", mock.Anything, newUser).Return(1, nil)

		userID, err := userService.RegisterUser(context.Background(), newUser)

		assert.NoError(t, err)
		assert.Equal(t, 1, userID)

		mockRepo.AssertExpectations(t)
	})

}

func TestLoginUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	t.Run("success login user", func(t *testing.T) {
		loginUser := &domain.User{Email: "john@example.com", Password: "password123"}
		expectedUser := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

		mockRepo.On("LoginUser", mock.Anything, loginUser).Return(expectedUser, nil)

		user, err := userService.LoginUser(context.Background(), loginUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.Email, user.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("fail login user not found error", func(t *testing.T) {
		loginUser := &domain.User{Email: "nonexistent@example.com", Password: "wrongpassword"}

		mockRepo.On("LoginUser", mock.Anything, loginUser).Return(&domain.User{ID: 0,
			Name: "", Email: "", Password: ""}, errors.New("no user found"))

		user, err := userService.LoginUser(context.Background(), loginUser)

		assert.Error(t, err)
		assert.Equal(t, user, &domain.User{ID: 0, Name: "", Email: "", Password: ""})
		assert.Equal(t, "no user found", err.Error())

		mockRepo.AssertExpectations(t)
	})
}
