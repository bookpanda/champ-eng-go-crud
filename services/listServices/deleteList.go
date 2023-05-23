package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func DeleteList(c *gin.Context) {
	id := c.Param("id")
	var list models.List
	database.DB.First(&list, id)
	database.DB.Delete(&models.List{}, id)

	c.JSON(200, gin.H{
		"list": list,
	})
}
