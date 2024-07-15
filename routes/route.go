package routes

import (
	"temp/controllers"

	"github.com/gin-gonic/gin"
)

// Initialize sets up routes for the application.
func Initialize(router *gin.Engine) {
	// Define routes for CRUD operations on products
	router.GET("/products", controllers.GetAllProducts)       // Get all products
	router.GET("/product/:id", controllers.GetProduct)        // Get a product by ID
	router.POST("/products", controllers.CreateProduct)       // Create a new product
	router.PUT("/products/:id", controllers.UpdateProduct)    // Update a product by ID
	router.DELETE("/products/:id", controllers.DeleteProduct) // Delete a product by ID
}
