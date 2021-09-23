// routes.go

package routes

import (
	"github.com/gin-gonic/gin"

	"REST-API/handlers"
)

func initializeRoutes() {

	router := gin.Default()

	// Handle the routes
	router.GET("/articles", handlers.getArticles)
	router.GET("/articles/:id", handlers.getArtByID)
	router.POST("/articles", handlers.postArticles)
	router.PUT("/articles/:id", handlers.putData)
}
