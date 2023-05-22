package routes

import "github.com/gin-gonic/gin"

func SetUpTaskRoutes(r *gin.Engine) {
	r.GET("/tasks")
}
