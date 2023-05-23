package utils

import (
	"errors"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"gorm.io/gorm"
)

func CheckListExists(id int) string {
	var list models.List
	if err := database.DB.Where("id = ?", id).First(&list).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "list with this id not found"
		}
		return "error creating a new task"
	}
	return ""
}
