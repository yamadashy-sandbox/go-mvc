package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (h HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home-index.tmpl", gin.H{
		"a": "a",
	})
}
