package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/pejota/bootstrap"
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

	// ...
}
