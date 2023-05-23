package services

import (
	"errors"

	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"gorm.io/gorm"
)

func CheckTaskExists(task *models.Task, id string) string {
	if err := database.DB.Where("id = ?", id).First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "task with this id not found"
		}
		return "error finding task"
	}
	return ""
}
