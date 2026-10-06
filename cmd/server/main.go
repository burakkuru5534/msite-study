package main

import (
	"context"
	"fmt"
	"github.com/burakkuru5534/msite-study/internal/user"
	"log"
)

func main() {
	repository := user.NewInMemoryRepository()
	service := user.NewService(repository)

	err := service.CreateUser(context.Background(), user.User{
		ID:   "1",
		Name: "Alice",
	})
	if err != nil {
		log.Fatal(err)
	}

	u, err := service.GetUser(context.Background(), "1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user: %+v\n", u)
}
