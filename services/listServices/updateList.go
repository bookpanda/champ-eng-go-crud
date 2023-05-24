package services

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

type UpdateListDto struct {
	Title string
	Order int
}

// UpdateList godoc
// @Summary      Update a List
// @Description  Update a List of ID with JSON. Only changes the fields that are in the JSON.
// @Tags         List
// @Accept json
// @Param id path string true "List ID"
// @Param ListDto body UpdateListDto true "UpdateListDto"
// @Produce      json
// @Router       /lists/{id} [put]
func UpdateList(c *gin.Context) {
	id := c.Param("id")
	body := UpdateListDto{}
	body.Order = -1
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var list models.List
	if res := CheckListExists(&list, id); res != "" {
		c.JSON(400, gin.H{
			"message": res,
		})
		return
	}
	if body.Title != "" {
		database.DB.Model(&list).Updates(models.List{Title: body.Title})
	}

	if body.Order != -1 {
		database.DB.Model(&list).Update("order", body.Order)
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}
