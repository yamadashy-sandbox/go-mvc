package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func (h *HelloController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}
