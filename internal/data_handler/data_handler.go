package data_handler

import (
	"strings"
        "text/tabwriter"
        "os"

	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DataHandler struct {
	Db_name         string
	Starting_weight float32
	Inc_weight      float32
	Decay_factor    float32
	Prune_threshold float32
}

func (data_h DataHandler) DropTable() {
	db, _ := sql.Open("sqlite3", data_h.Db_name)
	defer db.Close()
	statement, _ := db.Prepare("DROP TABLE IF EXISTS frecency;")
	statement.Exec()
}

func (data_h DataHandler) CreateTable() {
	db, _ := sql.Open("sqlite3", data_h.Db_name)
	defer db.Close()
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL);")
	statement.Exec()
}

func (data_h DataHandler) UpdatePath(path string) {
	db, _ := sql.Open("sqlite3", data_h.Db_name)
	defer db.Close()

	// check if path already exists
	existence_query := fmt.Sprintf(`SELECT COUNT(*) FROM frecency WHERE path="%s" LIMIT 1;`, path)
	row := db.QueryRow(existence_query)

	var path_exists_int uint
	row.Scan(&path_exists_int)

	var update_query string
	if path_exists_int == 1 {
		update_query = fmt.Sprintf(`UPDATE frecency SET score=score+%f
            WHERE path="%s";`, data_h.Inc_weight, path)
	} else {
		update_query = fmt.Sprintf(`INSERT INTO frecency VALUES ("%s", %f)`,
			path, data_h.Starting_weight)
	}
	statement, _ := db.Prepare(update_query)
	statement.Exec()

	// Decay the scores in the table
	decay_query := fmt.Sprintf(`UPDATE frecency SET score = %f*score;`, data_h.Decay_factor)
	statement, _ = db.Prepare(decay_query)
	statement.Exec()

	// Prune scores that are too low
	if data_h.Prune_threshold > 0 {
		query := fmt.Sprintf(`DELETE FROM frecency WHERE score < %f;`,
			data_h.Prune_threshold)

		statement, _ := db.Prepare(query)
		statement.Exec()
	}
}

func (data_h DataHandler) Size() uint {
	db, _ := sql.Open("sqlite3", data_h.Db_name)
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM frecency;")
	var count uint
	row.Scan(&count)

	return count
}

func (data_h DataHandler) ListPaths() string {
	db, _ := sql.Open("sqlite3", data_h.Db_name)
	defer db.Close()

	rows, _ := db.Query("SELECT path FROM frecency ORDER BY score;")
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

func (data_h DataHandler) PrintTable() {
	db, _ := sql.Open("sqlite3", data_h.Db_name)
	defer db.Close()

	rows, _ := db.Query("SELECT path, score FROM frecency ORDER BY score;")
	defer rows.Close()

	var path string
        var score float32
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	for rows.Next() {
		rows.Scan(&path, &score)
                tabbed_string := fmt.Sprintf("%.2f\t%s", score, path)
                fmt.Fprintln(writer, tabbed_string)
	}
        writer.Flush()
}
