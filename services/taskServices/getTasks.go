package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// GetTasks godoc
// @Summary      Get all Tasks
// @Description  Get all Tasks
// @Tags         Task
// @Produce      json
// @Router       /tasks [get]
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}
