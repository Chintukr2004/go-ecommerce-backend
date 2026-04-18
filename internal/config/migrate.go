package config

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("Migrations completed")
}
