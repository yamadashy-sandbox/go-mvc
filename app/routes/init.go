package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapRoutes() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	MapAPIRoutes(router)
	MapWebRoutes(router)

	http.Handle("/", router)
}
