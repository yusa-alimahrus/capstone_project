package services

import (
	"warehouse_inventory/database"

	"github.com/jinzhu/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		db: database.GetDB(),
	}
}

// Add your transaction service methods here
