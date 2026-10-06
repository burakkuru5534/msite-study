package user

import (
	"context"
	"errors"
)

var ErrInvalidUserID = errors.New("invalid user id")
var ErrInvalidUserName = errors.New("invalid user name")

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetUser(ctx context.Context, id string) (User, error) {
	if id == "" {
		return User{}, ErrInvalidUserID
	}

	return s.repository.Get(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, user User) error {
	if user.ID == "" {
		return ErrInvalidUserID
	}

	if user.Name == "" {
		return ErrInvalidUserName
	}

	return s.repository.Create(ctx, user)
}
