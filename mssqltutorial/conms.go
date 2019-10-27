package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "localhost\\SQLEXPRESS"
	port     = 1433
	user     = "sa"
	password = "password123"
	database = "tutorial"
)

var db *sql.DB
var err error

func main() {
	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err = sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	} else {
		fmt.Printf("Connected!\n")
	}
	defer db.Close()
	SelectVersion()
	// tsql := fmt.Sprintf("SELECT 1")
	// rows, err := db.Query(tsql)
	// if err != nil {
	// 	fmt.Println("Error reading rows: " + err.Error())

	// }
	// defer rows.Close()
}

func SelectVersion() {
	ctx := context.Background()

	err := db.PingContext(ctx)

	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	err = db.QueryRowContext(ctx, "select @@version").Scan(&result)

	if err != nil {
		log.Fatal("Scan Failed:", err.Error())
	}
	fmt.Printf("%s\n", result)

}
