package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title string
	Order int
	Tasks []Task `gorm:"constraint:OnDelete:CASCADE;"`
}
