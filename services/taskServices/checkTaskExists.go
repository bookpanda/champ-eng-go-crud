package services

import (
	"errors"
	"fmt"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"gorm.io/gorm"
)

func CheckTaskExists(task *models.Task, id string) string {
	if err := database.DB.Where("id = ?", id).First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Sprintf("Task with id of %s not found", id)
		}
		return fmt.Sprintf("Error finding Task with id of %s", id)
	}
	return ""
}
