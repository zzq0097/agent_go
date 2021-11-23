package api

import (
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LogLines(c *gin.Context) {
	log := c.Query("log")
	line, err := strconv.Atoi(c.Query("line"))
	if err != nil {
		line = 10
	}

	logs, err := service.ReadLogs(log, line)
	if err != nil {
		resp.FailMsg(err.Error(), c)
		return
	}
	resp.OkData(logs, c)
}
