package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the score of a path",
	Long:  `Update the score of a path`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Printf("Expected exactly one argument, got %d\n", len(args))
			os.Exit(1)
		}
		DataHandler.UpdatePath(args[0])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
