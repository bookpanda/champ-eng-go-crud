package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	id := c.Param("id")
	var list models.List
	database.DB.First(&list, id)

	c.JSON(200, gin.H{
		"list": list,
	})
}
