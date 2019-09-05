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
	"os/exec"
	"runtime"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var consumerApiKey string
var consumerApiSecret string

// authorizeCmd represents the authorize command
var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		consumerApiKey = viper.GetString("consumer_api_key")
		if consumerApiKey == "" {
			exitWithError(fmt.Errorf("consumer-api-key is required"))
		}

		consumerApiSecret = viper.GetString("consumer_api_secret")
		if consumerApiKey == "" {
			exitWithError(fmt.Errorf("consumer-api-secret is required"))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("authorize called")

		config := oauth1.Config{
			ConsumerKey:    consumerApiKey,
			ConsumerSecret: consumerApiSecret,
			Endpoint:       twitter.AuthorizeEndpoint,
			CallbackURL:    "oob",
		}

		requestToken, _, err := config.RequestToken()
		if err != nil {
			exitWithError(err)
		}

		authorizationURL, err := config.AuthorizationURL(requestToken)
		if err != nil {
			exitWithError(err)
		}

		authorizeUrl := authorizationURL.String()
		fmt.Printf("Open this URL in your browser:\n%s\n", authorizeUrl)
		err = openBrowser(authorizeUrl)
		if err != nil {
			fmt.Printf("open authorizationURL failed: %+v", err)
		}

		// receive
		fmt.Printf("Paste your PIN here: ")
		var verifier string
		_, err = fmt.Scanf("%s", &verifier)
		if err != nil {
			exitWithError(err)
		}
		accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)
		if err != nil {
			exitWithError(err)
		}

		viper.Set("access_token", accessToken)
		viper.Set("access_secret", accessSecret)

		err = viper.WriteConfig()
		if err != nil {
			exitWithError(err)
		}

		fmt.Println("authorize success")
	},
}

func init() {
	rootCmd.AddCommand(authorizeCmd)

	// Twitter Consumer API key
	authorizeCmd.Flags().String("consumer-api-key", "", "Twitter Consumer API Key")
	viper.BindPFlag("consumer_api_key", authorizeCmd.Flags().Lookup("consumer-api-key"))
	viper.BindEnv("consumer_api_key")

	// Twitter Consumer API Secret
	authorizeCmd.Flags().String("consumer-api-secret", "", "Twitter Consumer API Secret Key")
	viper.BindPFlag("consumer_api_secret", authorizeCmd.Flags().Lookup("consumer-api-secret"))
	viper.BindEnv("consumer_api_secret")
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
