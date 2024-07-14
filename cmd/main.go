package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/pejota/bootstrap"
	"github.com/marcelofabianov/pejota/internal/user"
	"github.com/marcelofabianov/pejota/internal/user/port"
	"github.com/marcelofabianov/pejota/pkg/postgres"
)

func main() {
	// Load config
	if err := bootstrap.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	config := bootstrap.NewConfig()
	ctx := context.Background()

	// Connect to database
	db, err := postgres.ConnectDB(*config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close(ctx)

	userContainer := user.NewUserContainer(db)

	err = userContainer.Invoke(func(service port.UserServiceApplication) {
		user, err := service.GetUser(port.GetUserInputServiceApplication{PublicID: "2401c307-ff9d-4963-895c-8cd7b97d9d67"})
		if err != nil {
			log.Fatalf("Error getting user: %v", err)
		}

		log.Printf("User: %+v", user)
	})

	if err != nil {
		log.Fatalf("Error invoking user container: %v", err)
	}
}
