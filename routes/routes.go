package routes

import (
	"warehouse_inventory/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Create controller instances
	categoryController := controllers.NewCategoryController()
	transactionController := controllers.NewTransactionController()

	// Category routes
	categoryRoutes := router.Group("/api/categories")
	{
		categoryRoutes.GET("", categoryController.GetCategories)
		categoryRoutes.GET("/:id", categoryController.GetCategory)
		categoryRoutes.POST("", categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	// Transaction routes
	transactionRoutes := router.Group("/api/transactions")
	{
		transactionRoutes.GET("", transactionController.GetTransactions)
		transactionRoutes.GET("/:id", transactionController.GetTransaction)
		transactionRoutes.POST("", transactionController.CreateTransaction)
		transactionRoutes.PUT("/:id", transactionController.UpdateTransaction)
		transactionRoutes.DELETE("/:id", transactionController.DeleteTransaction)
	}

	return router
}
