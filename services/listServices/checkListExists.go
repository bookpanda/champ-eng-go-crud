package services

import (
	"errors"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"gorm.io/gorm"
)

func CheckListExists(list models.List, id string) string {
	if err := database.DB.Where("id = ?", id).First(&list).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "list with this id not found"
		}
		return "error finding list"
	}
	return ""
}
