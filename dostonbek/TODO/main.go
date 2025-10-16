package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
  "time"
  "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// DBConfig holds the PostgreSQL connection credentials
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
func CreateToDoObjectHTTP(c *gin.Context, db *sql.DB) {
	var req ToDoObbject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	res, err := CreateToDoObject(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}

func GetToDoObjectsHTTP(c *gin.Context, db *sql.DB) {
	res, err := GetToDoObjects(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}
func main() {
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

	router := gin.Default()

	router.POST("/api/v1/create", func(c *gin.Context) {
		CreateToDoObjectHTTP(c, db)
	})

	router.GET("/api/v1/get-all", func(c *gin.Context) {
		GetToDoObjectsHTTP(c, db)
	})

	router.Run(":7000")
}

func CreateToDoObject(db *sql.DB, obj *ToDoObbject) (*ToDoObbject, error) {
	obj.CreatedAt = time.Now()
	query := `
		INSERT INTO tasks (title, created_at)
		VALUES ($1, $2)
		RETURNING id, title, created_at;
	`

	var result ToDoObbject
	res := db.QueryRow(query, obj.Title, obj.CreatedAt)

	if err := res.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if err := res.Scan(&result.ID, &result.Title, &result.CreatedAt); err != nil {
		log.Println("failed to parse")
		return nil, err
	}

	return &result, nil
}

func GetToDoObjects(db *sql.DB) ([]*ToDoObbject, error) {
	query := `
		SELECT id, title, created_at FROM tasks;
	`

	var result []*ToDoObbject

	res, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Close()

	// [1, 2, 3, 4, 5, 6]
	for res.Next() {
		var obj ToDoObbject
		if err := res.Scan(&obj.ID, &obj.Title, &obj.CreatedAt); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		result = append(result, &obj)
	}
	if err := res.Err(); err != nil {
		log.Println(err.Error())
	}

	return result, nil
}

type ToDoObbject struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}