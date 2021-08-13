package util

import (
	"agent/consts"
	"database/sql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", consts.DBFile)
	if err != nil {
		return nil
	}
	return db
}
