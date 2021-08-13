package route

import (
	"agent/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogRoute(r *gin.Engine) {
	log := r.Group("/log")
	{
		log.GET("/accessLog", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
		log.GET("/errorLog", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
	}
}
