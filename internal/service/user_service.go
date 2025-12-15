package service

import (
	"context"
	"errors"
	"inkspire/internal/model"
	"inkspire/internal/repository"
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters long")
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

/*
USER SERVICES ---
*/

func (s *UserService) CreateUser(ctx context.Context, email, password string) (*model.User, error) {
	if email == "" {
		return nil, ErrEmailRequired
	}

	if len([]rune(password)) < 8 {
		return nil, ErrPasswordTooShort
	}

	user := &model.User{
		Email:    email,
		Password: password, // hash LATER
	}

	return user, nil
}
