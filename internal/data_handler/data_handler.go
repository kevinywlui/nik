package data_handler

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() {
	database, _ := sql.Open("sqlite3", "./nik.db")
	statement, _ := database.Prepare(`
    CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);
    `)
	statement.Exec()
}
