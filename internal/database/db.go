package database

import (
	"api/internal/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDBURI())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to database")
	return db, nil
}
