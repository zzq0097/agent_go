package route

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommandRoute(r *gin.Engine) {
	command := r.Group("/command")
	{
		command.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxStatus(consts.NgxDefCmdFile, consts.NgxDefConfFile)))
		})
		command.GET("/start", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxStart(consts.NgxDefCmdFile, consts.NgxDefConfFile)))
		})
		command.GET("/stop", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxStop(consts.NgxDefCmdFile, consts.NgxDefConfFile)))
		})
		command.GET("/reload", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxReload(consts.NgxDefCmdFile, consts.NgxDefConfFile)))
		})
		command.GET("/ngxV", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxV(consts.NgxDefCmdFile)))
		})
		command.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxVersion(service.NgxV(consts.NgxDefCmdFile))))
		})
		command.GET("/cfgArgs", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxCfgArgs(service.NgxV(consts.NgxDefCmdFile))))
		})
	}
}
