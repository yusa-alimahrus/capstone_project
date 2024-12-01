package database

import (
	"log"
	"warehouse_inventory/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open("sqlite3", "warehouse_inventory.db")
	if err != nil {
		return nil, err
	}

	// Enable Logger
	DB.LogMode(true)

	// Auto Migrate
	DB.AutoMigrate(
		&models.Category{},
		// Add other models here
	)

	log.Println("SQLite database connected successfully")
	return DB, nil
}

func GetDB() *gorm.DB {
	return DB
}
