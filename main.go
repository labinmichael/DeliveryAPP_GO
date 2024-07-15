package main

import (
	"log"
	"os"
	"temp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode for better performance
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router with default middleware
	router := gin.New()

	// Use custom logger middleware for better logging
	router.Use(gin.Logger())

	// Recover middleware to handle panics and prevent crashes
	router.Use(gin.Recovery())

	// Initialize routes
	routes.Initialize(router)

	// Get the port from environment variables or default to :8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
