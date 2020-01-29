package data_handler

import (
	"strings"

	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DataHandler struct {
	db_name         string
	starting_weight float32
	inc_weight      float32
	decay_factor    float32
	prune_threshold float32
}

func (data_h DataHandler) DropTable() {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()
	statement, _ := db.Prepare("DROP TABLE IF EXISTS frecency;")
	statement.Exec()
}

func (data_h DataHandler) CreateTable() {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);")
	statement.Exec()
}

func (data_h DataHandler) UpdatePath(path string) {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	// check if path already exists
	existence_query := fmt.Sprintf(`SELECT COUNT(*) FROM frecency WHERE path="%s" LIMIT 1;`, path)
	row := db.QueryRow(existence_query)

	var path_exists_int uint
	row.Scan(&path_exists_int)

	var update_query string
	if path_exists_int == 1 {
		update_query = fmt.Sprintf(`UPDATE frecency SET score=score+%f
            WHERE path="%s";`, data_h.inc_weight, path)
	} else {
		update_query = fmt.Sprintf(`INSERT INTO frecency VALUES ("%s", %f)`,
			path, data_h.starting_weight)
	}
	statement, _ := db.Prepare(update_query)
	statement.Exec()
}

func (data_h DataHandler) Decay() {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	query := fmt.Sprintf(`UPDATE frecency SET score = %f*score;`, data_h.decay_factor)

	statement, _ := db.Prepare(query)
	statement.Exec()
}

func (data_h DataHandler) Prune() {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM frecency WHERE score < %f;`,
		data_h.prune_threshold)

	statement, _ := db.Prepare(query)
	statement.Exec()
}

func (data_h DataHandler) Size() uint {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM frecency;")
	var count uint
	row.Scan(&count)

	return count
}

func (data_h DataHandler) GetOrderedPaths() string {
	db, _ := sql.Open("sqlite3", data_h.db_name)
	defer db.Close()

	rows, _ := db.Query("SELECT path FROM frecency ORDER BY score LIMIT 10;")
	defer rows.Close()

	var sb strings.Builder
	var str string
	for rows.Next() {
		rows.Scan(&str)
		sb.WriteString(str)
		sb.WriteString("\n")
	}
	return sb.String()
}
