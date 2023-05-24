package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// DeleteTask godoc
// @Summary      Delete a Task
// @Description  Delete a Task with ID
// @Tags         Task
// @Accept json
// @Param id path string true "Task ID"
// @Produce      json
// @Router       /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if res := CheckTaskExists(&task, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}
	database.DB.Unscoped().Delete(&models.Task{}, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}
