package user

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)

	err := service.CreateUser(context.Background(), User{
		ID:   "1",
		Name: "Alice",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetUser(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)

	err := service.CreateUser(context.Background(), User{
		ID:   "1",
		Name: "Alice",
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	got, err := service.GetUser(context.Background(), "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID != "1" {
		t.Fatalf("expected ID 1, got %s", got.ID)
	}

	if got.Name != "Alice" {
		t.Fatalf("expected Alice, got %s", got.Name)
	}
}
