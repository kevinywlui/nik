package cmd

import (
	"fmt"
	"github.com/kevinywlui/nik/internal/data"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var DataHandler = data.DataHandler{
	DbFile:         "",
	IncWeight:      100,
	DecayFactor:    0.99,
	PruneThreshold: 1.0,
}

var rootCmd = &cobra.Command{
	Use:   "nik",
	Short: "cd using subsequence matching and frecency",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Set default data file
	var dotNikdotdb string
	dotNikdotdb, _ = homedir.Expand("~/.nik.db")
	DataHandler.DbFile = dotNikdotdb

	// Read environment variables for NIK_DBFILE, NIK_INCWEIGHT, NIK_DECAYFACTOR, NIK_PRUNETHRESHOLD
	if dbFileStr := os.Getenv("NIK_DBFILE"); len(dbFileStr) > 0 {
		DataHandler.DbFile = dbFileStr
	}
	if incWeightStr := os.Getenv("NIK_INCWEIGHT"); len(incWeightStr) > 0 {
		DataHandler.IncWeight, _ = strconv.ParseFloat(incWeightStr, 64)
	}
	if decayFactorStr := os.Getenv("NIK_DECAYFACTOR"); len(decayFactorStr) > 0 {
		DataHandler.DecayFactor, _ = strconv.ParseFloat(decayFactorStr, 64)
	}
	if pruneThresholdStr := os.Getenv("NIK_PRUNETHRESHOLD"); len(pruneThresholdStr) > 0 {
		DataHandler.PruneThreshold, _ = strconv.ParseFloat(pruneThresholdStr, 64)
	}
}
