package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Running")
	})
	CodeRouter(r)
	CommandRouter(r)
	LogRouter(r)
	ConfRouter(r)
	return r
}
