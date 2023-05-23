package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	database.DB.First(&task, id)
	database.DB.Delete(&models.Task{}, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}
