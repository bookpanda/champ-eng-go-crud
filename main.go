package main

import (
	"github.com/bookpanda/champ-eng-go-crud/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetUpTaskRoutes(r)
	routes.SetUpListRoutes(r)

	r.Run()
}
