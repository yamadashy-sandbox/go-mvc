package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	routing()
}

func routing() {
	router := gin.Default()

	routerApiV1 := router.Group("/api/v1")
	{
		routerApiV1.GET("/ping", pingAction)
	}

	http.Handle("/", router)
}

func pingAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
