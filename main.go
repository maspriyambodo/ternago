package main

import (
	"log"
	"ternago/backend-api/models"
	"ternago/backend-api/routes"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	
	// Create Gin router
	r := gin.Default()

	// Connect to the database
	models.ConnectDatabase()

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}
