package services

import (
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// GetTask godoc
// @Summary      Get a Task
// @Description  Get a Task with ID
// @Tags         Task
// @Produce      json
// @Param id path string true "Task ID"
// @Router       /tasks/{id} [get]
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if res := CheckTaskExists(&task, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}
