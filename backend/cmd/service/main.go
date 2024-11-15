package main

import (
	"log"

	"github.com/backend/database"
	"github.com/backend/internal/config"
)

func main() {
	cfg := config.Load("config/config.yaml")
	_, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
