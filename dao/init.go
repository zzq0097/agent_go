package dao

import (
	"agent/sql"
	"agent/util"
)

func CreateTable() {
	_, _ = util.DB.Exec(sql.CreateParameterTable)
}
