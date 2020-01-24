package data_handler

import (
	"testing"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func TestCreateTable(t *testing.T) {
	// create the table
	CreateTable()

	// check that it opens
	database, err := sql.Open("sqlite3", "./nik.db")
	if err != nil {
		t.Errorf("Unable to open database")
	}

	// check that the columns are correct
	rows, err := database.Query(`
    SELECT name FROM
    pragma_table_info('frecency');
    `)
	if err != nil {
		t.Errorf("Unable to open frecency database")
	}
	want := [2]string{"path", "score"}
	var got [2]string
	i := 0
	for rows.Next() {
		rows.Scan(&got[i])
		i++
	}
	if got != want {
		t.Errorf("Column mismatch, want %q, got %q", want, got)
	}
}
