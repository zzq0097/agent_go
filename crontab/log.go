package crontab

import (
	"agent/consts"
	"agent/model"
	"agent/service"
)

func CornReadLog() {
	sectionLog := model.SectionLog{}
	service.ReadOffset(consts.NgxRpmAccessLog, 0, &sectionLog)
}
