package main

import (
	"github.com/backend/database"
	"github.com/backend/internal/config"
	"log"
)

func main() {
	cfg := config.Load("config/config.yaml")
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()
	_, err = db.Db.Exec(`CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL
    );`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
}
