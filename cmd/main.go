package main

import (
	"log"

	"github.com/marcelofabianov/pejota.git/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}
