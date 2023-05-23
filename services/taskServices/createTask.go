package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/bookpanda/champ-eng-go-crud/utils"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var body struct {
		Description string
		DueDate     string
		Order       int
		ListID      int
	}
	c.Bind(&body)
	if res := utils.CheckListExists(body.ListID); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}
	task := models.Task{Description: body.Description, DueDate: body.DueDate, Order: body.Order, ListID: body.ListID}
	database.DB.Create(&task)

	c.JSON(200, gin.H{
		"task": task,
	})
}
