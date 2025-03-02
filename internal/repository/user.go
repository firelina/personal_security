package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"personal_security/internal/domain"
)

type UserRepositoryInterface interface {
	RegisterUser(ctx context.Context, newUser *domain.User) (int, error)
	LoginUser(ctx context.Context, loginUser *domain.User) (*domain.User, error)
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) RegisterUser(ctx context.Context, newUser *domain.User) (int, error) {
	query := "INSERT INTO personal_security.users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := r.pool.QueryRow(ctx, query, newUser.Name, newUser.Email, newUser.Password).Scan(&newUser.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to register user: %w", err)
	}
	return newUser.ID, nil
}

func (r *UserRepository) LoginUser(ctx context.Context, loginUser *domain.User) (*domain.User, error) {
	query := "SELECT id, email FROM  personal_security.users WHERE email = $1 AND password = $2"
	user := &domain.User{}
	err := r.pool.QueryRow(ctx, query, loginUser.Email, loginUser.Password).Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no user found")
		}
		return nil, fmt.Errorf("failed to login user: %w", err)
	}
	return user, nil
}
