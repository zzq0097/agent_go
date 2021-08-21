package route

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
)

func CommandRoute(r *gin.Engine) {
	command := r.Group("/command")
	{
		command.GET("/status", func(c *gin.Context) {
			resp.SuccessData(service.NgxStatus(consts.NgxDefCmdFile, consts.NgxDefConfFile), c)
		})
		command.GET("/start", func(c *gin.Context) {
			resp.SuccessData(service.NgxStart(consts.NgxDefCmdFile, consts.NgxDefConfFile), c)
		})
		command.GET("/stop", func(c *gin.Context) {
			resp.SuccessData(service.NgxStop(consts.NgxDefCmdFile, consts.NgxDefConfFile), c)
		})
		command.GET("/reload", func(c *gin.Context) {
			resp.SuccessData(service.NgxReload(consts.NgxDefCmdFile, consts.NgxDefConfFile), c)
		})
		command.GET("/ngxV", func(c *gin.Context) {
			resp.SuccessData(service.NgxV(consts.NgxDefCmdFile), c)
		})
		command.GET("/version", func(c *gin.Context) {
			resp.SuccessData(service.NgxVersion(service.NgxV(consts.NgxDefCmdFile)), c)
		})
		command.GET("/cfgArgs", func(c *gin.Context) {
			resp.SuccessData(service.NgxCfgArgs(service.NgxV(consts.NgxDefCmdFile)), c)
		})
	}
}
