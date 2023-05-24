package main

import (
	"log"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/bookpanda/champ-eng-go-crud/utils"
)

func init() {
	utils.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	if database.DB.Migrator().HasTable(&models.List{}) {
		database.DB.Migrator().DropTable(&models.List{})
	}
	if database.DB.Migrator().HasTable(&models.Task{}) {
		database.DB.Migrator().DropTable(&models.Task{})
	}
	database.DB.AutoMigrate(&models.Task{})
	database.DB.AutoMigrate(&models.List{})
	log.Println("Migrate tables successful")
}
