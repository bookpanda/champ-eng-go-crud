package routes

import (
	"github.com/bookpanda/champ-eng-go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpListRoutes(r *gin.Engine) {
	r.GET("/lists", controllers.GetLists)
	r.GET("/lists/:id", controllers.GetList)
	r.POST("/lists", controllers.CreateList)
	r.PUT("/lists/:id", controllers.UpdateList)
	r.DELETE("/lists/:id", controllers.DeleteList)
}
