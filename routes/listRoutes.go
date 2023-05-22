package routes

import "github.com/gin-gonic/gin"

func SetUpListRoutes(r *gin.Engine) {
	r.GET("/lists")
}
