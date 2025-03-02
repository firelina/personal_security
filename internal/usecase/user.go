package usecase

import (
	"context"
	"personal_security/internal/domain"
	"personal_security/internal/repository"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) RegisterUser(ctx context.Context, newUser *domain.User) (int, error) {
	resultChan := make(chan struct {
		ID  int
		Err error
	})

	go func() {
		userID, err := u.repo.RegisterUser(ctx, newUser)
		resultChan <- struct {
			ID  int
			Err error
		}{ID: userID, Err: err}
	}()

	result := <-resultChan
	return result.ID, result.Err
}

func (u *UserService) LoginUser(ctx context.Context, loginUser *domain.User) (*domain.User, error) {
	resultChan := make(chan struct {
		User *domain.User
		Err  error
	})

	go func() {
		user, err := u.repo.LoginUser(ctx, loginUser)
		resultChan <- struct {
			User *domain.User
			Err  error
		}{User: user, Err: err}
	}()

	result := <-resultChan
	return result.User, result.Err
}
