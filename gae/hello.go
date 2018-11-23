package gae

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	router := gin.Default()

	router.GET("/ping", pingAction)

	http.Handle("/", router)
}

func pingAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
