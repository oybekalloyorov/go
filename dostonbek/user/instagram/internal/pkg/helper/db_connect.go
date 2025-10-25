package helper

import (
	"database/sql"
	"fmt"
	"instagram/internal/pkg/config"
)

func ConnectPostgress(cfg *config.DBConfig) (*sql.DB, error) {
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
