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
	database.DB.AutoMigrate(&models.Task{})
	database.DB.AutoMigrate(&models.List{})
	log.Println("Migrate tables successful")
}
