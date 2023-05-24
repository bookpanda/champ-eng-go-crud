package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

type CreateListDto struct {
	Title string
	Order int
}

// CreateList godoc
// @Summary      Create a new List
// @Description  Create a new List from JSON
// @Tags         lists
// @Accept json
// @Param ListDto body CreateListDto true "CreateListDto"
// @Produce      json
// @Router       /lists [post]
func CreateList(c *gin.Context) {
	body := CreateListDto{}
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
	list := models.List{Title: body.Title, Order: body.Order}
	database.DB.Create(&list)

	c.JSON(200, gin.H{
		"list": list,
	})
}
