package data

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/kevinywlui/nik/internal/util"
	_ "github.com/mattn/go-sqlite3"
	homedir "github.com/mitchellh/go-homedir"
	"os"
	"strings"
	"text/tabwriter"
)

type DataHandler struct {
	Db_file         string
	Starting_weight float32
	Inc_weight      float32
	Decay_factor    float32
	Prune_threshold float32
}

// isFile returns whether path exists on the filesystem
func isFile(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// DropTable will drop the frecency table if it exists. An error will be
// returned if `data_h.Db_file` is not found.
func (data_h DataHandler) DropTable() error {
	// Check if `data_h.Db_file` exists
	if !isFile(data_h.Db_file) {
		return fmt.Errorf("Database file %s not found", data_h.Db_file)
	}

	// Open the database
	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()

	_, err := db.Exec("DROP TABLE IF EXISTS frecency;")
	if err != nil {
		return fmt.Errorf("Unable to execute the drop query")
	}

	return nil
}

// CreateTable will create a sql table at `data_h.Db_file`. This will create
// the file if it does not exists.
func (data_h DataHandler) CreateTable() error {
	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS frecency (path TEXT PRIMARY KEY, score REAL, prefix TEXT, base TEXT);")
	if err != nil {
		return fmt.Errorf("Unable to execute the select query")
	}
	return nil
}

// UpdatePath updates the table given a `path`. This should increase the score
// of `path` and decrease the score of all other paths.
//
// The home directory will be ignored.
func (data_h DataHandler) UpdatePath(path string) {
	// Ignore the home directory
	home_str, _ := homedir.Dir()
	if path == home_str {
		return
	}

	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()

	// check if path already exists
	existence_query := fmt.Sprintf(`SELECT COUNT(*) FROM frecency WHERE path=%q LIMIT 1;`, path)
	row := db.QueryRow(existence_query)

	var path_exists_int uint
	row.Scan(&path_exists_int)

	// Increment score if path already exists. Insert into table otherwise.
	var update_query string
	if path_exists_int == 1 {
		update_query = fmt.Sprintf(`UPDATE frecency SET score=score+%f
            WHERE path="%s";`, data_h.Inc_weight, path)

	} else {
		// Compute the tokens if needed.
		tokens := util.PrefixBaseSplit(path)
		prefix := tokens[0]
		base := tokens[1]
		update_query = fmt.Sprintf(`INSERT INTO frecency VALUES (%q, %f, %q, %q)`,
			path, data_h.Starting_weight, prefix, base)
	}
	db.Exec(update_query)

	// Decay the scores in the table
	decay_query := fmt.Sprintf(`UPDATE frecency SET score = %f*score;`, data_h.Decay_factor)
	db.Exec(decay_query)

	// Prune scores that are too low
	if data_h.Prune_threshold > 0 {
		prune_query := fmt.Sprintf(`DELETE FROM frecency WHERE score < %f;`,
			data_h.Prune_threshold)

		db.Exec(prune_query)

	}
}

// Size returns the number of rows to the table
func (data_h DataHandler) Size() uint {
	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM frecency;")
	var count uint
	row.Scan(&count)

	return count
}

// ListPaths list the paths ordered using their scores and the `descending` flag.
func (data_h DataHandler) ListPaths(descending bool, scores bool) (string, error) {
	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()

	var rows *sql.Rows
	var err error
	if scores {
		if descending {
			rows, err = db.Query("SELECT path, score FROM frecency ORDER BY score DESC;")
		} else {
			rows, err = db.Query("SELECT path, score FROM frecency ORDER BY score;")
		}
		defer rows.Close()
		if err != nil {
			return "", fmt.Errorf("Unable to execute select")
		}

		var path string
		var score float32

		// Use a bytes buffer to later return a string
		buf := new(bytes.Buffer)
		writer := tabwriter.NewWriter(buf, 0, 8, 1, '\t', tabwriter.AlignRight)

		for rows.Next() {
			rows.Scan(&path, &score)
			tabbed_string := fmt.Sprintf("%.2f\t%s", score, path)
			fmt.Fprintln(writer, tabbed_string)
		}
		writer.Flush()

		return buf.String(), nil
	} else {
		if descending {
			rows, err = db.Query("SELECT path FROM frecency ORDER BY score DESC;")
		} else {
			rows, err = db.Query("SELECT path FROM frecency ORDER BY score;")
		}
		defer rows.Close()
		if err != nil {
			return "", fmt.Errorf("Unable to execute select")
		}

		var sb strings.Builder
		var str string
		for rows.Next() {
			rows.Scan(&str)
			sb.WriteString(str)
			sb.WriteString("\n")
		}
		return sb.String(), nil
	}
}

// GetTopBaseMatch will return the matching path with the highest score.
func (data_h DataHandler) GetTopBaseMatch(query_base string) (string, bool, error) {
	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()

	rows, err := db.Query("SELECT path, base FROM frecency ORDER BY score DESC;")
	defer rows.Close()
	if err != nil {
		return "", false, fmt.Errorf("Unable to execute select")
	}

	var path, base string
	for rows.Next() {
		rows.Scan(&path, &base)
		if util.IsSubsequence(query_base, base) {
			return path, true, nil
		}
	}
	return "", false, nil
}

func (data_h DataHandler) GetTopPrefixBaseMatch(query_prefix, query_base string) (string, bool, error) {
	db, _ := sql.Open("sqlite3", data_h.Db_file)
	defer db.Close()

	rows, err := db.Query("SELECT path, prefix, base FROM frecency ORDER BY score DESC;")
	defer rows.Close()
	if err != nil {
		return "", false, fmt.Errorf("Unable to execute select")
	}

	var path, prefix, base string
	for rows.Next() {
		rows.Scan(&path, &prefix, &base)
		if util.IsSubsequence(query_prefix, prefix) && util.IsSubsequence(query_base, base) {
			return path, true, nil
		}
	}
	return "", false, nil
}
