package routes

import (
	"kiokumushi-api/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapAPIRoutes() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	hello := new(controllers.HelloController)

	router.GET("/health", health.Status)

	routerAPIV1 := router.Group("/api/v1")
	{
		routerAPIV1.GET("/ping", hello.Ping)
	}

	http.Handle("/", router)
}
