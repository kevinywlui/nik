package data_handler

import (
	"testing"
)

// use a test db
var test_db = "./test.db"

func TestFull(t *testing.T) {
	test_data_h := DataHandler{test_db}

        // drop any old table
	test_data_h.DropTable()
	// create the table
	test_data_h.CreateTable()

	// add some paths
	test_data_h.AddPath("A", 10)
	test_data_h.AddPath("B", 100)

	// verify that there are 2 paths
	var want_pre_decay uint
	want_pre_decay = 2
	got_pre_decay := test_data_h.Size()
	if got_pre_decay != want_pre_decay {
		t.Errorf("Incorrect number of rows, want %d, got %d", want_pre_decay, got_pre_decay)
	}

	// decay the table and prune the A path away
	// verify that only B is left
	test_data_h.Decay(0.5)
	test_data_h.Prune(10)

	var want_post_decay uint
	want_post_decay = 1
	got_post_decay := test_data_h.Size()
	if got_post_decay != want_post_decay {
		t.Errorf("Incorrect number of rows, want %d, got %d", want_post_decay, got_post_decay)
	}
}
