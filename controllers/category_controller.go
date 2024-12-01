package controllers

import (
	"net/http"
	"strconv"
	"warehouse_inventory/models"
	"warehouse_inventory/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		service: services.NewCategoryService(),
	}
}

func (c *CategoryController) GetCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) GetCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	category, err := c.service.GetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.ID = uint(id)

	if err := c.service.UpdateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.service.DeleteCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
