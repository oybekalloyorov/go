package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectPostgress(ctg DBConfig)(*sql.DB, error){
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		ctg.Host, ctg.Port, ctg.User, ctg.Password, ctg.DBName, ctg.SSLMode,
	)

	db, err := sql.Open("postgres", dns)
	if (err != nil) {
		return nil, fmt.Errorf("Failed to open connection: %w", err)
	}

	 // Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func CreateToDoObjectHTTP(c *gin.Context, db *sql.DB){
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

func GetToDoObjectsHTTP(c *gin.Context, db *sql.DB){
	res, err := GetToDoObjects(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res })
}

// ✅ GetByID handler
func GetToDoByIDHTTP(c *gin.Context, db *sql.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := GetToDoByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}

// ✅ DeleteByID handler
func DeleteToDoByIDHTTP(c *gin.Context, db *sql.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rows, err := DeleteToDoByID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no record deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
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

	db, err := ConnectPostgress(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	defer db.Close()

	fmt.Println("Successfullu connected to PostgreSQL!")

	router := gin.Default()

	router.POST("/api/v1/create", func(c *gin.Context) {
		CreateToDoObjectHTTP(c, db)
	})

	router.GET("/api/v1/get-all", func(c *gin.Context) {
		GetToDoObjectsHTTP(c, db)
	})

	router.GET("/api/v1/get/:id", func(c *gin.Context) {
		GetToDoByIDHTTP(c, db)
	})

	router.DELETE("/api/v1/delete/:id", func(c *gin.Context) {
		DeleteToDoByIDHTTP(c, db)
	})

	router.Run(":7000")

}

func CreateToDoObject(db *sql.DB, obj *ToDoObbject)(*ToDoObbject, error){
	obj.CreatedAt = time.Now()

	query := `
		INSERT INTO taskstest(title, created_at)
		VALUES ($1, $2)
		RETURNING id, title, created_at;
	`

	var result ToDoObbject

	res := db.QueryRow(query, obj.Title, obj.CreatedAt)

	if err := res.Scan(&result.ID, &result.Title, &result.CreatedAt); err != nil {
		log.Panicln("Failed to parse")
		return nil, err
	}

	return &result, nil
}

func GetToDoObjects(db *sql.DB) ([]*ToDoObbject, error){
	query := `
		SELECT id, title, created_at from taskstest;
	`

	var result []*ToDoObbject

	res, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
		return nil, err
	}

	defer res.Close()

	for res.Next(){
		var obj ToDoObbject
		if err := res.Scan(&obj.ID, &obj.Title, &obj.CreatedAt); err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}
		result = append(result, &obj)
	}
	if err := res.Err(); err != nil {
		log.Panicln(err.Error())
	}
	return result, nil
}
// ✅ GetByID function
func GetToDoByID(db *sql.DB, id int) (*ToDoObbject, error) {
	query := `SELECT id, title, created_at FROM taskstest WHERE id=$1;`

	var obj ToDoObbject
	err := db.QueryRow(query, id).Scan(&obj.ID, &obj.Title, &obj.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ✅ DeleteByID function
func DeleteToDoByID(db *sql.DB, id int) (int64, error) {
	query := `DELETE FROM taskstest WHERE id=$1;`

	res, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
type ToDoObbject struct{
	ID int `json:"id"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}