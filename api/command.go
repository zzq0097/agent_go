package api

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
)

func NgxStatus(c *gin.Context) {
	ngxStatus := service.NgxStatus(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.SuccessData(ngxStatus, c)
}

func NgxStart(c *gin.Context) {
	ngxStatus := service.NgxStart(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.SuccessData(ngxStatus, c)
}

func NgxReload(c *gin.Context) {
	ngxStatus := service.NgxReload(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.SuccessData(ngxStatus, c)
}

func NgxStop(c *gin.Context) {
	ngxStatus := service.NgxStop(consts.NgxDefCmdFile, consts.NgxDefConfFile)
	resp.SuccessData(ngxStatus, c)
}

func NgxV(c *gin.Context) {
	resp.SuccessData(service.NgxV(consts.NgxDefCmdFile), c)
}

func NgxVersion(c *gin.Context) {
	resp.SuccessData(service.NgxVersion(service.NgxV(consts.NgxDefCmdFile)), c)
}

func NgxCfgArgs(c *gin.Context) {
	resp.SuccessData(service.NgxCfgArgs(service.NgxV(consts.NgxDefCmdFile)), c)
}
