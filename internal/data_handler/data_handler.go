package data_handler

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db_name string) {
	database, _ := sql.Open("sqlite3", db_name)
	statement, _ := database.Prepare(`
    CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);
    `)
	statement.Exec()
}
