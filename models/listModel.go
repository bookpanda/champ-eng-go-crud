package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title string
	Order int
	Tasks []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ListID"`
}
