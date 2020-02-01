package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
        "github.com/kevinywlui/nik/internal/config_nik"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the paths in the table ordered by their frecency score",
	Long: `List the ptahs in the table ordered by their frecency score`,
	Run: func(cmd *cobra.Command, args []string) {
                dh := config_nik.DataHandler
                print_score, _ := cmd.Flags().GetBool("score")
                if print_score {
                    dh.PrintTable()
                } else {
                    fmt.Printf(dh.ListPaths())
                }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
        listCmd.Flags().BoolP("score", "s", false, "print the scores")
}
