package user

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")

type Repository interface {
	Get(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) error
}

type InMemoryRepository struct {
	users map[string]User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[string]User),
	}
}

func (r *InMemoryRepository) Get(ctx context.Context, id string) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryRepository) Create(ctx context.Context, user User) error {
	if _, exists := r.users[user.ID]; exists {
		return ErrUserAlreadyExists
	}

	r.users[user.ID] = user

	return nil
}
