package dao

import (
	"agent/consts"
	"agent/util"
	"fmt"
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

func InsertSectionLog(json string) int64 {
	fmt.Println(json)
	if result, err := util.DB.Exec("insert into ngx_log values (?,?)", util.Now(), json); err == nil {
		if affected, err := result.RowsAffected(); err == nil {
			return affected
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	return 0
}
