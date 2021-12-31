package dao

import (
	"agent/consts"
	"agent/util"
	"strconv"
)

func GetLogOffset() (int64, error) {
	row := util.DB.QueryRow("select `value` from `parameter` where `name` = ?", consts.AccessLogOffset)
	var offset string
	err := row.Scan(&offset)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(offset, 10, 64)
}

func InsertSectionLog(jsonStr string) int64 {
	if result, err := util.DB.Exec("insert into ngx_log values (?,?)", util.Now(), jsonStr); err == nil {
		if affected, err := result.RowsAffected(); err == nil {
			return affected
		}
	}
	return 0
}
