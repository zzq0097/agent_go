package route

import (
	"agent/resp"
	"github.com/gin-gonic/gin"
)

func CodeRoute(r *gin.Engine) {
	code := r.Group("/code")
	{
		code.GET("/get", func(c *gin.Context) {
			resp.Success(c)
		})
	}
}
