/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/kevinywlui/nik/internal/config_nik"
	"github.com/spf13/cobra"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the highest scoring path matching this string",
	RunE: func(cmd *cobra.Command, args []string) error {
		if n_args := len(args); n_args != 1 {
			return fmt.Errorf("There should be exactly 1 argument, got %d", n_args)
		}
		// small will be the string to be matched
		small := args[0]

		dh := config_nik.DataHandler
		dh.GetTopMatch(small)
		path, found, err := dh.GetTopMatch(small)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
