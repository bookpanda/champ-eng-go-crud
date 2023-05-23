package services

import (
	"strconv"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/bookpanda/champ-eng-go-crud/utils"
	"github.com/gin-gonic/gin"
)

func MoveTaskToAnotherList(c *gin.Context) {
	taskID := c.Param("taskID")
	listID, err := strconv.ParseInt(c.Param("listID"), 10, 0)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "list id is not number",
		})
		return
	}

	if res := utils.CheckListExists(int(listID)); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}
	var task models.Task
	database.DB.First(&task, taskID)
	database.DB.Model(&task).Update("list_id", listID)
	// database.DB.Model(&task).Updates(models.Task{Description: body.Description, DueDate: body.DueDate, Order: body.Order, ListID: body.ListID})

	c.JSON(200, gin.H{
		"task": task,
	})
}
