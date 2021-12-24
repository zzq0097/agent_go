package crontab

import (
	"agent/consts"
	"agent/service"
)

func CornReadLog() {
	service.ReadOffset(consts.NgxRpmAccessLog, 0, service.LogLine)
}
