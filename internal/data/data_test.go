package data

import (
	"fmt"
	"testing"
)

func TestFull(t *testing.T) {
	dh := DataHandler
	dh.GetTopMatch("kevin")
}
