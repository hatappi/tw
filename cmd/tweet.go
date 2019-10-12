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

	"github.com/hatappi/tw/editor"
	"github.com/hatappi/tw/twitter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "tweet message",
	Long:  "tweet message",
	Run: func(cmd *cobra.Command, args []string) {
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			exitWithError(err)
		}
		if message == "" {
			txt, editErr := editor.EditText()
			if editErr != nil {
				exitWithError(editErr)
			}
			message = string(txt)
		}

		if viper.GetBool("dry-run") {
			fmt.Print(message)
			return
		}

		config, err := twitter.LoadConfigFromViper()
		if err != nil {
			exitWithError(err)
		}
		client := twitter.NewClient(config)

		p := &twitter.UpdateStatusParams{
			Status: message,
		}
		t, err := client.StatusService.UpdateStatus(p)
		if err != nil {
			exitWithError(err)
		}
		fmt.Printf("%s\nby %s\n", t.Text, t.User.Name)
	},
}

func init() {
	tweetCmd.Flags().StringP("message", "m", "", "the message to send")

	tweetCmd.Flags().Bool("dry-run", false, "dry run. print message instead of send")
	err := viper.BindPFlag("dry-run", tweetCmd.Flags().Lookup("dry-run"))
	if err != nil {
		exitWithError(err)
	}
	err = viper.BindEnv("dry-run", "DRY_RUN")
	if err != nil {
		exitWithError(err)
	}

	rootCmd.AddCommand(tweetCmd)
}
