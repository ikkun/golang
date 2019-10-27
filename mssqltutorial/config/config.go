package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB(db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=127.0.0.1\\SQLEXPRESS;user id=sa;password=password123;port=1433;database=tutorial")
	return
}
