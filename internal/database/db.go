package database

import (
	"api/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDBURI())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	err = createInitialSchema(db)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to database")
	return db, nil
}

func Disconnect(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
		return
	}
}

func createInitialSchema(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS	 tasks (
		id SERIAL PRIMARY KEY,
		taskid TEXT NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		priority INT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	log.Printf("Successfully created tasks table in database")
	return nil
}
