package dao

import (
	"agent/util"
)

func InsertNgxStatus(connections int, requests int) int64 {
	if result, err := util.DB.Exec("insert into ngx_status values (?,?,?)", util.Now(), connections, requests); err == nil {
		if affected, err := result.RowsAffected(); err == nil {
			return affected
		}
	}
	return 0
}
