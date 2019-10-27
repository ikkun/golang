package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB
var err error

type auth_user struct {
	gorm.Model
	username string
	email    string
}

var (
	server   = "localhost\\SQLEXPRESS"
	port     = 1433
	user     = "sa"
	password = "password123"
	database = "tutorial"
)

func InitialMigration() {
	// connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
	// 	user, password, server, port, database)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	db, err := gorm.Open("mssql", connString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
}
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Endpoint Hit")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	db, err := gorm.Open("mssql", connString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	var users []auth_user
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
