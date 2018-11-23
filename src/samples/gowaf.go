package samples

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitNetHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
}

func InitGin() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!!")
	})
	http.Handle("/", r)
}
