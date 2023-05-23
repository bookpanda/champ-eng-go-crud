package main

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/routes"
	"github.com/bookpanda/champ-eng-go-crud/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.SetUpTaskRoutes(r)
	routes.SetUpListRoutes(r)

	r.Run()
}
