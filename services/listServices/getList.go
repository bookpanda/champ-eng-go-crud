package services

import (
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	id := c.Param("id")
	var list models.List
	if res := CheckListExists(&list, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}
