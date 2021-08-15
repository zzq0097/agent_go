package route

import (
	"github.com/gin-gonic/gin"
)

func MainRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Running")
	})
	CodeRoute(r)
	CommandRoute(r)
	LogRoute(r)
	ConfRoute(r)
	return r
}
