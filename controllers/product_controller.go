package controllers

import (
	"net/http"
	"temp/models"

	"github.com/gin-gonic/gin"
)

// GetAllProducts retrieves all products.
func GetAllProducts(c *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct retrieves a single product by ID.
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := models.GetProductByID(id)
	if err != nil {
		// Check if the error indicates that the product was not found
		if err == models.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		// Handle other internal server errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product.
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := models.CreateProduct(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// UpdateProduct updates an existing product.
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct models.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.UpdateProduct(id, updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct deletes an existing product.
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
