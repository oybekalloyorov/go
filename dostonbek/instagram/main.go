package main

import (
	"database/sql"
	"fmt"
	"log"
)
type DBConfig struct {
  Host     string
  Port     int
  User     string
  Password string
  DBName   string
  SSLMode  string
}

// ConnectPostgres establishes a connection to PostgreSQL and returns *sql.DB
func ConnectPostgres(cfg DBConfig) (*sql.DB, error) {
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

  return db, nil
}
func main(){
cfg := DBConfig{
    Host:     "localhost",
    Port:     5432,
    User:     "postgres",
    Password: "oybek",
    DBName:   "todo_app",
    SSLMode:  "disable",
  }

  db, err := ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to PostgreSQL!")
}