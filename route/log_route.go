package route

import (
	"agent/resp"
	"github.com/gin-gonic/gin"
)

func LogRoute(r *gin.Engine) {
	log := r.Group("/log")
	{
		log.GET("/accessLog", func(c *gin.Context) {
			resp.Success(c)
		})
		log.GET("/errorLog", func(c *gin.Context) {
			resp.Success(c)
		})
	}
}
