package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Description string
		DueDate     string
	}
	c.Bind(&body)
	var task models.Task
	database.DB.First(&task, id)
	database.DB.Model(&task).Updates(models.Task{Description: body.Description, DueDate: body.DueDate})

	c.JSON(200, gin.H{
		"task": task,
	})
}
