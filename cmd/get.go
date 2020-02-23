package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the highest scoring path matching this string",
	RunE: func(cmd *cobra.Command, args []string) error {
		n_args := len(args)
		if n_args != 1 && n_args != 2 {
			return fmt.Errorf("There should be exactly 1 or 2 argument, got %d", n_args)
		}
		// Use either just base matching or prefix-base matching
		// depending on the number of arguments
		var path string
		var found bool
		var err error
		if n_args == 1 {
			base_query := args[0]
			path, found, err = DataHandler.GetTopBaseMatch(base_query)
		} else {
			prefix_query := args[0]
			base_query := args[1]
			path, found, err = DataHandler.GetTopPrefixBaseMatch(prefix_query, base_query)
		}

		if err != nil {
			return err
		}
		if !found {
			os.Exit(1)
		}
		fmt.Println(path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
