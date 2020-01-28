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
	defer db.Close()
	statement, _ := db.Prepare(`
CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);
    `)
	statement.Exec()

}

func (data_h DataHandler) AddPath(path string, starting_weight float32) {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	query := fmt.Sprintf(`
INSERT INTO frecency VALUES ("%s", %f);
	`, path, starting_weight)

	statement, _ := db.Prepare(query)
	statement.Exec()

}

func (data_h DataHandler) Decay(decay_factor float32) {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	query := fmt.Sprintf(`
UPDATE frecency SET score = %f*score;
`, decay_factor)

	statement, _ := db.Prepare(query)
	statement.Exec()

}

func (data_h DataHandler) Prune(threshold float32) {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	query := fmt.Sprintf(`
DELETE FROM frecency WHERE score < %f;
`, threshold)

	statement, _ := db.Prepare(query)
	statement.Exec()

}

func (data_h DataHandler) Size() uint {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()
	rows, _ := db.Query(`
SELECT COUNT(*) FROM frecency;
    `)
	defer rows.Close()
	rows.Next()
	var count uint
	rows.Scan(&count)

	return count
}
