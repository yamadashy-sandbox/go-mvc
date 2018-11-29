package routes

import (
	"kiokumushi-api/app/controllers/api"

	"github.com/gin-gonic/gin"
)

func MapAPIRoutes(router *gin.Engine) {
	// Controller
	helloController := new(api.HelloController)

	// Routes
	routerAPIV1 := router.Group("/api/v1")
	{
		routerAPIV1.GET("/hello", helloController.Index)
	}
}
