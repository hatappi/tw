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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// greetingCmd represents the greeting command
var greetingCmd = &cobra.Command{
	Use:   "greeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		parentMessage, err := cmd.Parent().PersistentFlags().GetString("message")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if parentMessage != "" {
			fmt.Printf("ParentCommand message is %s\n", parentMessage)
		}

		msg := "Hello!!"
		if viper.GetString("greeting") != "" {
			msg = viper.GetString("greeting")
		}
		fmt.Printf("%s\n", msg)

		viper.Set("test", "test test")
		err = viper.WriteConfig()
		if err != nil {
			panic(err)
		}
		fmt.Printf("save file path is %s\n", viper.ConfigFileUsed())
	},
}

type Hoge struct {
	Name string
}

func init() {
	echoCmd.AddCommand(greetingCmd)

	// 1. { "greeting": "おはようございます" } in ~/.tw.json
	// 2. TW_GREETING=hello go run main.go echo greeting
	// 3. go run main.go echo greeting --msg hello!!!!
	greetingCmd.Flags().String("msg", "", "greeting msg")

	err := viper.BindPFlag("greeting", greetingCmd.Flags().Lookup("msg"))
	if err != nil {
		exitWithError(err)
	}

	err = viper.BindEnv("greeting", "TW_GREETING")
	if err != nil {
		exitWithError(err)
	}

}
