package main

import (
	"log"

	"./configs"
	"./database"
	"./models"
	"./repositories"
)

func main() {
	//database configs
	dbServer, dbUser, dbPassword, dbName := "localhost\\SQLEXPRESS", "sa", "password123", "tutorial"
	db, err := database.ConnectToDB(dbServer, dbUser, dbPassword, dbName)

	// unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}

	//ping to database
	err = db.DB().Ping()

	//error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	//migration
	db.AutoMigrate(&models.Contact{})

	defer db.Close()

	contactRepository := repositories.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8081")
}
