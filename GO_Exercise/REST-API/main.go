package main

import (
	"REST-API/routes"
)

func main() {

	// Initialize the routes
	r := routes.InitializeRoutes()

	// Start serving the application
	r.Run("localhost:5000")
}
