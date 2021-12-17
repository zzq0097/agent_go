package util

import (
	"database/sql"
	"os"
	"path/filepath"
)

var Home, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func InitDB() *sql.DB {
	dbFile := Home + "sqlite.db"
	if !FileExist(dbFile) {
		_, _ = os.Create(dbFile)
	}
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil
	}
	return db
}
