package route

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfRoute(r *gin.Engine) {
	conf := r.Group("/conf")
	{
		conf.GET("/allFile", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.AllConf(consts.NgxDefConf)))
		})
		conf.GET("/replace", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.AllConf(consts.NgxDefConf)))
		})
		conf.GET("/rollback", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.AllConf(consts.NgxDefConf)))
		})

	}
}
