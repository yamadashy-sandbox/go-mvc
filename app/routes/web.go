package routes

import (
	"kiokumushi-api/app/controllers"

	"github.com/gin-gonic/gin"
)

func MapWebRoutes(router *gin.Engine) {
	// Controller
	homeController := new(controllers.HomeController)
	healthController := new(controllers.HealthController)

	// Routes
	router.GET("/", homeController.Index)
	router.GET("/health", healthController.Status)

	// http.Handle("/", router)
}
