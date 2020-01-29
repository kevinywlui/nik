package data_handler

import (
	"testing"
)

// use a test db
var test_db = "./test.db"

func TestFull(t *testing.T) {
	test_data_h := DataHandler{
            db_name: test_db,
            starting_weight: 100,
            inc_weight: 10,
            decay_factor: 0.5,
            prune_threshold: 55.0,
	}

	// drop any old table
	test_data_h.DropTable()
	// create the table
	test_data_h.CreateTable()

	// add some paths
	test_data_h.UpdatePath("A")
	test_data_h.UpdatePath("A")
	test_data_h.UpdatePath("A")
	test_data_h.UpdatePath("A")
	test_data_h.UpdatePath("B")

	// verify that there are 2 paths
	var want_pre_decay uint
	want_pre_decay = 2
	got_pre_decay := test_data_h.Size()
	if got_pre_decay != want_pre_decay {
		t.Errorf("Incorrect number of rows, want %d, got %d", want_pre_decay, got_pre_decay)
	}

	// decay the table and prune the A path away
	// verify that only B is left
	test_data_h.Decay()
	test_data_h.Prune()

	var want_post_decay uint
	want_post_decay = 1
	got_post_decay := test_data_h.Size()
	if got_post_decay != want_post_decay {
		t.Errorf("Incorrect number of rows, want %d, got %d", want_post_decay, got_post_decay)
	}
}
