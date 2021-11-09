package main

import (
	db "USRMGMT/controllers"
	"USRMGMT/routes"
	"fmt"

	"github.com/gin-contrib/cors"
)

func main() {

	db.InitDatabase()

	// Create a new table.
	// Initialize the routes
	r := routes.InitializeRoutes()

	r.Use(cors.Default())

	// Start the application server.
	fmt.Println("Server is running on port 5000")
	r.Run("localhost:5000")
}
