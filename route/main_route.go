package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Running")
	})
	CodeRoute(r)
	CommandRoute(r)
	LogRoute(r)
	ConfRoute(r)
	return r
}
