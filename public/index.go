package public

import (
	"net/http"

	"kiokumushi-api/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.MapAPIRoutes(router)
	routes.MapWebRoutes(router)

	http.Handle("/", router)
}
