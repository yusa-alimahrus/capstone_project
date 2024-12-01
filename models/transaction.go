package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	ProductID       uint   `json:"product_id"`
	Quantity        int    `json:"quantity"`
	TransactionType string `json:"transaction_type"` // 'in' or 'out'
}
