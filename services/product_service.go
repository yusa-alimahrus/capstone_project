package services

import (
	"warehouse_inventory/database"

	"github.com/jinzhu/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService() *ProductService {
	return &ProductService{
		db: database.GetDB(),
	}
}

// Add your product service methods here
