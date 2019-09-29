package cmd

import (
	"fmt"

	"github.com/hatappi/tw/twitter"
	"github.com/spf13/cobra"
)

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use:   "timeline",
	Short: "timeline command",
	Long:  "timeline command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var timelineHomeCmd = &cobra.Command{
	Use:   "home",
	Short: "show Home Timeline tweet",
	Long:  "show Home Timeline tweet",
	Run: func(cmd *cobra.Command, args []string) {
		config := twitter.LoadConfigFromViper()
		client := twitter.NewClient(config)

		cnt, err := cmd.Flags().GetInt("count")
		if err != nil {
			exitWithError(err)
		}

		params := &twitter.HomeTimelineParams{
			Count: cnt,
		}
		tweets, err := client.TimelineService.GetHomeTimeline(params)
		if err != nil {
			exitWithError(err)
		}

		for _, t := range tweets {
			fmt.Printf("%s(@%s)\n%s\n--------------\n", t.User.Name, t.User.ScreenName, t.Text)
		}
	},
}

func init() {
	timelineHomeCmd.Flags().IntP("count", "c", 0, "tweet count")
	timelineCmd.AddCommand(timelineHomeCmd)

	rootCmd.AddCommand(timelineCmd)
}
