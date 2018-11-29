package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
