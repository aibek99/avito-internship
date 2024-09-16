package main

import (
	"context"
	"log"

	"avito-internship/internal/config"
	"avito-internship/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("godotenv.Load: %w", err)
	}
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to get all configs")
	}

	source := server.NewServer(configs)

	err = source.Run(context.Background())
	if err != nil {
		log.Printf("source.Run: %w", err)
	}
}
