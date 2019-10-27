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
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")
	defer db.Close()
	createID, err := CreateAuthor("I", "IKKUN", "1981-12-31")
	fmt.Printf("Inserted ID: %d successfully.\n", createID)
	// SelectVersion()
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

func CreateAuthor(lang string, author string, oob string) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		log.Fatal("What?")
	}

	//Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("insert into programming_programming_authors values(@lang,@author,@oob)")

	// Execute non-query with named parametes
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("programming_languages", lang),
	)
	if err != nil {
		log.Fatal("Error inserting new row: " + err.Error())
		return -1, err
	}
	return 0, 1

}
