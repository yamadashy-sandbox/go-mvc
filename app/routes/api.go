package routes

import (
	"kiokumushi-api/app/controllers"

	"github.com/gin-gonic/gin"
)

func MapAPIRoutes(router *gin.Engine) {
	// Controller
	helloController := new(controllers.HelloController)

	// Routes
	routerAPIV1 := router.Group("/api/v1")
	{
		routerAPIV1.GET("/ping", helloController.Status)
	}
}
