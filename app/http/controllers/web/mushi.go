package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MushiController struct{}

func (h MushiController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "mushi-index.tmpl", gin.H{
		"a": "a",
	})
}
