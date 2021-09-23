// routes.go

package routes

import (
	"github.com/gin-gonic/gin"

	"REST-API/handlers"
)

func InitializeRoutes() {

	router := gin.Default()

	// Handle the routes
	router.GET("/articles", handlers.GetArticles)
	router.GET("/articles/:id", handlers.GetArtByID)
	router.POST("/articles", handlers.PostArticles)
	router.PUT("/articles/:id", handlers.PutData)
}
