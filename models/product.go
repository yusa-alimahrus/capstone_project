package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name       string  `json:"name"`
	CategoryID uint    `json:"category_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}
