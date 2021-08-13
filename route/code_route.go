package route

import (
	"agent/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CodeRoute(r *gin.Engine) {
	code := r.Group("/code")
	{
		code.GET("/get", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
	}
}
