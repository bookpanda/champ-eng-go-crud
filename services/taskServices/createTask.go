package services

import (
	"strconv"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	services "github.com/bookpanda/champ-eng-go-crud/services/listServices"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var body struct {
		Description string
		DueDate     string
		Order       int
		ListID      int
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if body.Description == "" || body.DueDate == "" {
		c.JSON(400, gin.H{
			"message": "fields description or duedate is empty",
		})
		return
	}
	var list models.List
	if res := services.CheckListExists(list, strconv.Itoa(body.ListID)); res != "" {
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
