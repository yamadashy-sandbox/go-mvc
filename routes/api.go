package routes

import (
	"github.com/stivan622/kiokumushi-api/app/http/controllers/api"

	"github.com/gin-gonic/gin"
)

func MapAPIRoutes(router *gin.Engine) {
	// Controller
	helloController := &api.HelloController{}

	// Routes
	routerAPIV1 := router.Group("/api/v1")
	{
		routerAPIV1.GET("/hello", helloController.Index)
	}
}
