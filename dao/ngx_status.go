package dao

import (
	"agent/util"
)

func InsertNgxStatus(connections int, requests int) {
	_, _ = util.DB.Exec("insert into ngx_status values (?,?,?)", util.Now(), connections, requests)
}
