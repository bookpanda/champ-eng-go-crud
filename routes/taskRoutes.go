package routes

import (
	"github.com/bookpanda/champ-eng-go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpTaskRoutes(r *gin.Engine) {
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTask)
	r.POST("/tasks", controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}
