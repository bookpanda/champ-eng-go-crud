package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// DeleteList godoc
// @Summary      Delete a List
// @Description  Delete a List with ID, and its Tasks
// @Tags         List
// @Accept json
// @Param id path string true "List ID"
// @Produce      json
// @Router       /lists/{id} [delete]
func DeleteList(c *gin.Context) {
	id := c.Param("id")
	var list models.List
	if res := CheckListExists(&list, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}
	database.DB.Unscoped().Delete(&models.List{}, id)

	c.JSON(200, gin.H{
		"list": list,
	})
}
