package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func ReorderTask(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Order int
	}
	c.Bind(&body)
	var task models.Task
	database.DB.First(&task, id)
	database.DB.Model(&task).Update("order", body.Order)
	// database.DB.Model(&task).Updates(models.Task{Description: body.Description, DueDate: body.DueDate, Order: body.Order, ListID: body.ListID})

	c.JSON(200, gin.H{
		"task": task,
	})
}
