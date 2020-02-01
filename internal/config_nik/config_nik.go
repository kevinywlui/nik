package config_nik

import (
        "github.com/kevinywlui/nik/internal/data_handler"
)

var DataHandler = data_handler.DataHandler{
                    Db_name: "./nik.db",
                    Starting_weight: 100,
                    Inc_weight: 10,
                    Decay_factor: 0.5,
                    Prune_threshold: 1.0,
                }

func Configure() {}
