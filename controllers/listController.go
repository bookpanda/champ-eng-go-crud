package controllers

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetLists(c *gin.Context) {
	var lists []models.List
	database.DB.Find(&lists)

	c.JSON(200, gin.H{
		"lists": lists,
	})
}

func GetList(c *gin.Context) {
	id := c.Param("id")
	var list models.List
	database.DB.First(&list, id)

	c.JSON(200, gin.H{
		"list": list,
	})
}

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

func DeleteList(c *gin.Context) {
	id := c.Param("id")
	var list models.List
	database.DB.First(&list, id)
	database.DB.Delete(&models.List{}, id)

	c.JSON(200, gin.H{
		"list": list,
	})
}
