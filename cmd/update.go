package cmd

import (
        "github.com/kevinywlui/nik/internal/config_nik"
	"github.com/spf13/cobra"
        "os"
        "fmt"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the score of a path",
	Long: `Update the score of a path`,
	Run: func(cmd *cobra.Command, args []string) {
                dh := config_nik.DataHandler
                if len(args) != 1 {
                    fmt.Printf("Expected exactly one argument, got %d\n", len(args))
                    os.Exit(1)
                }
                dh.UpdatePath(args[0])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
