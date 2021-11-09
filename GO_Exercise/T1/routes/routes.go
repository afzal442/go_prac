package routes

import (
	"USRMGMT/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/view-profile", controllers.ViewProfile)
	router.GET("/logout", controllers.Logout)

	return router
}
