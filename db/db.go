package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// ConnectDB establishes a connection to the database using the connection URL
func ConnectDB() (*sql.DB, error) {
	// Get the connection URL from the environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verify the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
