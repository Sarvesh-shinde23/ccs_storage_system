package main

import (
	"github.com/Sarvesh-shinde23/ccs_storage_system/controllers" // Importing controller package for handling API endpoints

	"github.com/gin-contrib/cors" // Importing cors middleware for Cross-Origin Resource Sharing
	"github.com/gin-gonic/gin"    // Importing Gin for creating HTTP server
)

// main is the entry point of the program.
func main() {
	// Create a new Gin router with default settings.
	r := gin.Default()

	// Enable CORS (Cross-Origin Resource Sharing) middleware and allow all origins.
	r.Use(cors.Default())

	// Define API endpoints
	// Endpoint for creating a file or folder
	r.POST("/api/create", controllers.CreateFileOrFolder)

	// Run the server on port 8080
	r.Run(":8080")
}
