package cmd

import (
	"fmt"
	"github.com/kevinywlui/nik/internal/config_nik"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the paths in the table ordered by their frecency score",
	Long:  `List the ptahs in the table ordered by their frecency score`,
	Run: func(cmd *cobra.Command, args []string) {
		dh := config_nik.DataHandler
		scores, _ := cmd.Flags().GetBool("scores")
		descending, _ := cmd.Flags().GetBool("descending")
		str, err := dh.ListPaths(descending, scores)
		if err != nil {
			fmt.Println("Error listing table")
		} else {
			fmt.Print(str)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("scores", "s", false, "print the scores")
	listCmd.Flags().BoolP("descending", "d", false, "print in descending order of scores")
}
