package database

import (
	"database/sql"
	"fmt"

	"github.com/backend/internal/config"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Db *sql.DB
}

func NewDB(cfg *config.ConfigDB) (*DataBase, error) {
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.HostDB,
		cfg.PortDB,
		cfg.NameDB,
	)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DataBase{Db: db}, nil
}
