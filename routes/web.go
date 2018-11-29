package routes

import (
	"kiokumushi-api/app/http/controllers/web"

	"github.com/gin-gonic/gin"
)

func MapWebRoutes(router *gin.Engine) {
	// Controller
	homeController := new(web.HomeController)
	healthController := new(web.HealthController)

	// Routes
	router.GET("/", homeController.Index)
	router.GET("/health", healthController.Index)

	// http.Handle("/", router)
}
