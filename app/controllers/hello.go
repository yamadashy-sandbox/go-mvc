package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func (h HelloController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}