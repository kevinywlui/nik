package data_handler

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DataHandler struct {
	db_name string
}

func (data_h DataHandler) CreateTable() {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	statement, _ := db.Prepare(`
CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);
    `)
	statement.Exec()

	defer db.Close()
}

func (data_h DataHandler) AddPath(path string) {
	db, _ := sql.Open("sqlite3", data_h.db_name)

	const default_weight = 10.0
	query := fmt.Sprintf(`
INSERT INTO frecency VALUES ("%s", %f);
	`, path, default_weight)

	statement, _ := db.Prepare(query)
	statement.Exec()

	defer db.Close()
}

func (data_h DataHandler) DecayTable() {
	db, _ := sql.Open("sqlite3", data_h.db_name)

	const decay_factor = 0.99
	query := fmt.Sprintf(`
UPDATE frecency SET score = %f*score;
`, decay_factor)

	statement, _ := db.Prepare(query)
	statement.Exec()

	defer db.Close()
}
