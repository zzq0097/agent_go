package route

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
)

func ConfRoute(r *gin.Engine) {
	conf := r.Group("/conf")
	{
		conf.GET("/allFile", func(c *gin.Context) {
			resp.SuccessData(service.AllConf(consts.NgxDefConf), c)
		})
		conf.GET("/replace", func(c *gin.Context) {
			resp.SuccessData(service.AllConf(consts.NgxDefConf), c)
		})
		conf.GET("/rollback", func(c *gin.Context) {
			resp.SuccessData(service.AllConf(consts.NgxDefConf), c)
		})
	}
}
