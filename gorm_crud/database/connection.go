package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func ConnectToDB(dbServer string, dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	// var connectionString = fmt.Sprintf(
	// 	"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	dbUser, dbPassword, dbName,
	// )

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=1433;database=%s;",
		dbServer, dbUser, dbPassword, dbName)
	return gorm.Open("mssql", connString)

	// return gorm.Open("mysql", connectionString)
}
