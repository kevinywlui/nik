package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		DataHandler.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
