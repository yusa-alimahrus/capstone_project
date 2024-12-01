package services

import (
	"warehouse_inventory/database"
	"warehouse_inventory/models"

	"github.com/jinzhu/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		db: database.GetDB(),
	}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Find(&categories).Error
	return categories, err
}

func (s *CategoryService) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := s.db.First(&category, id).Error
	return category, err
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.db.Create(category).Error
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return s.db.Save(category).Error
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.db.Delete(&models.Category{}, id).Error
}
