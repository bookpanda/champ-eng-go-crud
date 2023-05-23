package services

import (
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if res := CheckTaskExists(task, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}
