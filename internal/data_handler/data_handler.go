package data_handler

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DataHandler struct {
	db_name string
}

func CreateTable(data_h DataHandler) {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	statement, _ := db.Prepare(`
    CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);
    `)
	statement.Exec()

	defer db.Close()
}

func DecayTable(db_name string) {
	// loop through all rows and multiply scores by 0.99
	const decay_factor = 0.99
}
