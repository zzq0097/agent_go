package dao

import (
	"agent/util"
)

func GetParameter(name string) (string, error) {
	row := util.DB.QueryRow("select `value` from `parameter` where `name` = ? limit 1", name)
	var value string
	return value, row.Scan(&value)
}

func SaveOrUpdateParameter(name string, value string) int64 {
	if result, err := util.DB.Exec("replace into `parameter` values (?,?)", name, value); err == nil {
		if affected, err := result.RowsAffected(); err == nil {
			return affected
		}
	}
	return 0
}

func SaveOrIgnoreParameter(name string, value string) int64 {
	if result, err := util.DB.Exec("insert or ignore into `parameter` values (?,?)", name, value); err == nil {
		if affected, err := result.RowsAffected(); err == nil {
			return affected
		}
	}
	return 0
}

func batch(list map[string]string) {
	if tx, err := util.DB.Begin(); err == nil {
		for name, value := range list {
			_, _ = tx.Exec("replace into `parameter` values (?,?)", name, value)
		}
		_ = tx.Commit()
	}
}
