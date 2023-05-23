package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func UpdateList(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title string
		Order int
	}
	c.Bind(&body)
	var list models.List
	database.DB.First(&list, id)
	database.DB.Model(&list).Updates(models.List{Title: body.Title, Order: body.Order})

	c.JSON(200, gin.H{
		"list": list,
	})
}
