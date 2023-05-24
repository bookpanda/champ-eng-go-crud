package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

// DeleteList godoc
// @Summary      Delete a List
// @Description  Delete a List with ID, and its Tasks
// @Tags         lists
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

	// uintID, err := strconv.ParseUint(id, 10, 64)
	// if err != nil {
	// 	fmt.Printf("Parse Error: %s", err)
	// }

	// err = database.DB.Model(&models.List{
	// 	Model: gorm.Model{
	// 		ID: uint(uintID),
	// 	},
	// }).Association("Tasks").Delete()
	// if err != nil {
	// 	fmt.Printf("Error Log: %s", err)
	// }
	// database.DB.Unscoped().Select(clause.Associations).Unscoped().Delete(&models.List{}, id)

	database.DB.Delete(&models.List{}, id)
	database.DB.Where("list_id", id).Delete(&models.Task{})
	// database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&list)
	// database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&models.List{}, id)
	// database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&models.Task{}, id)

	c.JSON(200, gin.H{
		"list": list,
	})
}
