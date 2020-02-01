package cmd

import (
        "github.com/kevinywlui/nik/internal/config_nik"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initial the database",
	Long: `Initialize the database`,
	Run: func(cmd *cobra.Command, args []string) {
                dh := config_nik.DataHandler
                dh.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
