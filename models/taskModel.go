package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	DueDate     string
	Order       int
	ListID      int
}
