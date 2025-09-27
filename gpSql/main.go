package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "oybek"
	dbname   = "first_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `INSERT INTO "Employee"("Name", "EmpId") VALUES('Rohit', 21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)
	// fmt.Println("Inserted 1 row successfully")
	inserDynStmt := `INSERT INTO "Employee"("Name", "EmpId") VALUES($1, $2)`
	_, e = db.Exec(inserDynStmt, "Kirish", 03)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}