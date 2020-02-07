package config_nik

import (
	"testing"
)

func TestFull(t *testing.T) {
	dh := DataHandler
	dh.GetTopMatch("kevin")
}
