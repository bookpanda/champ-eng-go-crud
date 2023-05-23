package services

import (
	"strconv"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	services "github.com/bookpanda/champ-eng-go-crud/services/listServices"
	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
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
	var task models.Task
	if res := CheckTaskExists(task, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
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
	database.DB.Model(&task).Updates(models.Task{Description: body.Description, DueDate: body.DueDate, Order: body.Order, ListID: body.ListID})

	c.JSON(200, gin.H{
		"task": task,
	})
}
