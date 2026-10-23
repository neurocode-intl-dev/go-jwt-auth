package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectPostgres(databaseURL string) error {
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Test the connection
	if err := db.Ping(context.Background()); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	DB = db
	log.Println("✓ Successfully connected to PostgreSQL database")
	return nil
}

// CreateUserTable creates the users table if it doesn't exist
func CreateUserTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to create users table: %w", err)
	}

	log.Println("✓ Users table created/verified successfully")
	return nil
}
