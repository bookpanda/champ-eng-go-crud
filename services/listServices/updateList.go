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
	body.Order = -1
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if body.Title == "" {
		c.JSON(400, gin.H{
			"message": "fields title is empty",
		})
		return
	}
	var list models.List
	if res := CheckListExists(list, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}
	database.DB.Model(&list).Updates(models.List{Title: body.Title})

	if body.Order != -1 {
		database.DB.Model(&list).Updates(models.List{Order: body.Order})
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}
