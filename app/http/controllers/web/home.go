package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (h HomeController) Index(c *gin.Context) {
	c.String(http.StatusOK, "トップページ")
}
