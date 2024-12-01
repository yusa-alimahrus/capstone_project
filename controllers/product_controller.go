package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get products"})
}

func AddProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Add product"})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update product"})
}

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete product"})
}
