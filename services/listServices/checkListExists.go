package services

import (
	"errors"
	"fmt"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"gorm.io/gorm"
)

func CheckListExists(list *models.List, id string) string {
	if err := database.DB.Where("id = ?", id).First(list).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Sprintf("List with id of %s not found", id)
		}
		return fmt.Sprintf("Error finding List with id of %s", id)
	}
	return ""
}
