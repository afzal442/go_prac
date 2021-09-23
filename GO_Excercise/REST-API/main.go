package main

import (
	"github.com/gin-gonic/gin"

	"REST-API/routes"
)

func main() {
	router := gin.Default()

	// Initialize the routes
	routes.InitializeRoutes()

	// Start serving the application
	router.Run()
}
