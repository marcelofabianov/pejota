package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/pejota.git/bootstrap"
	"github.com/marcelofabianov/pejota.git/pkg/postgres"
)

func main() {
	// Load config
	if err := bootstrap.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx := context.Background()

	// Connect to database
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close(ctx)
}
