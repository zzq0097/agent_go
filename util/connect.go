package util

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var Home, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var DB *sql.DB

func InitDB() {
	dbFile := Home + "/sqlite.db"
	if !FileExist(dbFile) {
		_, err := os.Create(dbFile)
		if err != nil {
			return
		}
	}
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	DB = db
}
