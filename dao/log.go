package dao

import (
	"agent/consts"
	"agent/model"
	"agent/service"
	"agent/util"
	"encoding/json"
	"strconv"
)

func GetLogOffset() (int64, error) {
	query, err := util.DB.Query("select `value` from `parameter` where `name` = ?", consts.AccessLogOffset)
	if err != nil {
		return 0, err
	}
	var offset string
	err = query.Scan(offset)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(offset, 10, 64)
}

func InsertSectionLog() {
	offset, err := GetLogOffset()
	if err != nil {
		return
	}
	sectionLog := model.SectionLog{}
	service.ReadOffset(consts.NgxRpmAccessLog, offset, &sectionLog)
	marshal, err := json.Marshal(sectionLog)
	if err != nil {
		return
	}
	_, _ = util.DB.Exec("insert into ngx_log values (?,?)", util.Now(), string(marshal))
}
