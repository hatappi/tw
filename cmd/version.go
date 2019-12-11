/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version",
	Run: func(cmd *cobra.Command, args []string) {
		short, err := cmd.Flags().GetBool("short")
		if err != nil {
			exitWithError(err)
		}

		if version == "" {
			version = "None"
		}
		if commit == "" {
			commit = "None"
		}

		if short {
			fmt.Println(version)
		} else {
			fmt.Printf("version is %s, commit is %s\n", version, commit)
		}
	},
}

func init() {
	versionCmd.Flags().BoolP("short", "s", false, "show short version")
	rootCmd.AddCommand(versionCmd)
}
