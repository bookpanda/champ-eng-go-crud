package services

import (
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// GetList godoc
// @Summary      Get a List
// @Description  Get a List with ID
// @Tags         List
// @Produce      json
// @Param id path string true "List ID"
// @Router       /lists/{id} [get]
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
