package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// GetLists godoc
// @Summary      Get all Lists
// @Description  Get all Lists
// @Tags         List
// @Produce      json
// @Router       /lists [get]
func GetLists(c *gin.Context) {
	var lists []models.List
	database.DB.Find(&lists)

	c.JSON(200, gin.H{
		"lists": lists,
	})
}
