package api

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
)

func NgxStatus(c *gin.Context) {
	ngxStatus := service.NgxStatus(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.OkData(ngxStatus, c)
}

func NgxStart(c *gin.Context) {
	ngxStatus := service.NgxStart(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.OkData(ngxStatus, c)
}

func NgxReload(c *gin.Context) {
	ngxStatus := service.NgxReload(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.OkData(ngxStatus, c)
}

func NgxStop(c *gin.Context) {
	ngxStatus := service.NgxStop(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.OkData(ngxStatus, c)
}

func NgxV(c *gin.Context) {
	resp.OkData(service.NgxV(consts.NgxDefCmdFile), c)
}

func NgxVersion(c *gin.Context) {
	resp.OkData(service.NgxVersion(service.NgxV(consts.NgxDefCmdFile)), c)
}

func NgxCfgArgs(c *gin.Context) {
	resp.OkData(service.NgxCfgArgs(service.NgxV(consts.NgxDefCmdFile)), c)
}

func NgxTConf(c *gin.Context) {
	resp.OkData(service.NgxTConf(consts.NgxDefCmdFile, consts.NgxDefConfFile), c)
}

func CheckPortUsed(c *gin.Context) {
	resp.OkData(service.CheckPortUsed(), c)
}
