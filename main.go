package main

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/docs"
	"github.com/bookpanda/champ-eng-go-crud/routes"
	"github.com/bookpanda/champ-eng-go-crud/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utils.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.SetUpTaskRoutes(r)
	routes.SetUpListRoutes(r)

	docs.SwaggerInfo.Title = "User API documentation"
	docs.SwaggerInfo.Description = "Basic RESTful CRUD for Backend Assignement"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
