package main

import (
	"github.com/backend/database"
	"github.com/backend/internal/config"
	"log"
)

func main() {
	cfg := config.Load("config/config.yaml")
	_, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
