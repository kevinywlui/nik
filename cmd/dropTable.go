package cmd

import (
        "github.com/kevinywlui/nik/internal/config_nik" 
	"github.com/spf13/cobra"
)

// dropTableCmd represents the dropTable command
var dropTableCmd = &cobra.Command{
	Use:   "dropTable",
	Short: "Drop the frecency table",
	Long: `Drop the frecency table`,
	Run: func(cmd *cobra.Command, args []string) {
            dh := config_nik.DataHandler
            dh.DropTable()
	},
}

func init() {
	rootCmd.AddCommand(dropTableCmd)
}
