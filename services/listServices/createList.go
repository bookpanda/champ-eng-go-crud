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
	c.Bind(&body)
	list := models.List{Title: body.Title, Order: body.Order}
	database.DB.Create(&list)

	c.JSON(200, gin.H{
		"list": list,
	})
}
