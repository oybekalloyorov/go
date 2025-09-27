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

	// insertStmt := `INSERT INTO "Employee"("Name", "EmpId") VALUES('Rohit', 21)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)
	// // fmt.Println("Inserted 1 row successfully")
	// inserDynStmt := `INSERT INTO "Employee"("Name", "EmpId") VALUES($1, $2)`
	// _, e = db.Exec(inserDynStmt, "Kirish", 03)
	// CheckError(e)

	createStmt := `
	CREATE TABLE IF NOT EXISTS "TestTable" (
		id SERIAL PRIMARY KEY,
		"Name" TEXT NOT NULL,
		"EmpId" INT NOT NULL UNIQUE
	);
	`
	if _, err := db.Exec(createStmt); err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")

	//insert some data
	insertDynStmt := `INSERT INTO "TestTable"("Name", "EmpId") VALUES($1, $2)`
	_, e := db.Exec(insertDynStmt, "Alice", 101)
	CheckError(e)
	_, e = db.Exec(insertDynStmt, "Bob", 102)
	CheckError(e)
	_, e = db.Exec(insertDynStmt, "Charlie", 103)
	CheckError(e)
	fmt.Println("Inserted 3 rows successfully")
	//query some data
	rows, err := db.Query(`SELECT * FROM "TestTable"`)
	CheckError(err)
	defer rows.Close()

	// Delete some data
	deleteStmt := `DELETE FROM "TestTable" WHERE "EmpId"=$1`
	_, err = db.Exec(deleteStmt, 103)
	CheckError(err)
	fmt.Println("Deleted 1 row successfully")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}