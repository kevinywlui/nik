package config_nik

import (
	"github.com/kevinywlui/nik/internal/data"
	homedir "github.com/mitchellh/go-homedir"
)

var DataHandler = data.DataHandler{
	Db_file:         "nik.db",
	Inc_weight:      100,
	Decay_factor:    0.99,
	Prune_threshold: 1.0,
}

func init() {
	p := &DataHandler
	path := "~/.nik.db"
	p.Db_file, _ = homedir.Expand(path)
}
