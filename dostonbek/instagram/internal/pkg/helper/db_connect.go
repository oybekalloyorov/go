package helper

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"oybekalloyorov/salom/dostonbek/instagram/internal/pkg/config"
)

// ConnectPostgres establishes a connection to PostgreSQL and returns *sql.DB
func ConnectPostgres(cfg *config.DBConfig) (*sql.DB, error) {
	// Build connection string
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// Open connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Connected to PostgreSQL successfully")
	return db, nil
}
