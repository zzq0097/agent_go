package api

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
)

func AccessLog(c *gin.Context) {
	logs, err := service.ReadLogs(consts.NgxRpmAccessLog, 10)
	if err != nil {
		resp.FailMsg(err.Error(), c)
	}
	resp.OkData(logs, c)
}

func ErrorLog(c *gin.Context) {
	logs, err := service.ReadLogs(consts.NgxDefErrorLog, 10)
	if err != nil {
		resp.FailMsg(err.Error(), c)
	}
	resp.OkData(logs, c)
}
