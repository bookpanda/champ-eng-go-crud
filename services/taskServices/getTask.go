package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	database.DB.First(&task, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}
