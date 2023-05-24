package services

import (
	"strconv"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	services "github.com/bookpanda/champ-eng-go-crud/services/listServices"
	"github.com/gin-gonic/gin"
)

type UpdateTaskDto struct {
	Description string
	DueDate     string
	Order       int
	ListID      int
}

// UpdateTask godoc
// @Summary      Update a Task
// @Description  Update a Task of ID with JSON. Only changes the fields that are in the JSON.
// @Tags         Task
// @Accept json
// @Param id path string true "Task ID"
// @Param TaskDto body UpdateTaskDto true "UpdateTaskDto"
// @Produce      json
// @Router       /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTaskDto{}
	body.Order = -1
	body.ListID = -1
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var task models.Task
	if res := CheckTaskExists(&task, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}

	if body.Description != "" {
		database.DB.Model(&task).Updates(models.Task{Description: body.Description})
	}
	if body.DueDate != "" {
		database.DB.Model(&task).Updates(models.Task{DueDate: body.DueDate})
	}
	if body.Order != -1 {
		database.DB.Model(&task).Update("order", body.Order)
	}
	if body.ListID != -1 {
		var list models.List
		if res := services.CheckListExists(&list, strconv.Itoa(body.ListID)); res != "" {
			c.JSON(400, gin.H{
				"message": res,
			})
			return
		}
		database.DB.Model(&task).Update("list_id", body.ListID)
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}
