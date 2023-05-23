package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var body struct {
		Title string
		Order int
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	list := models.List{Title: body.Title, Order: body.Order}
	database.DB.Create(&list)

	c.JSON(200, gin.H{
		"list": list,
	})
}
