package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text"`
}
