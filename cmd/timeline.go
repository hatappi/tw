package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/hatappi/tw/twitter"
	"github.com/spf13/cobra"
)

var (
	twitterClient *twitter.Client
	isFollow      bool
)

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use:   "timeline",
	Short: "timeline command",
	Long:  "timeline command",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config, err := twitter.LoadConfigFromViper()
		if err != nil {
			exitWithError(err)
		}
		twitterClient = twitter.NewClient(config)

		isFollow, err = cmd.Parent().PersistentFlags().GetBool("follow")
		if err != nil {
			exitWithError(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			exitWithError(err)
		}
	},
}

var timelineHomeCmd = &cobra.Command{
	Use:   "home",
	Short: "show Home Timeline tweet",
	Long:  "show Home Timeline tweet",
	Run: func(cmd *cobra.Command, args []string) {
		cnt, err := cmd.Flags().GetInt("count")
		if err != nil {
			exitWithError(err)
		}
		isExcludeReplies, err := cmd.Flags().GetBool("exclude-replies")
		if err != nil {
			exitWithError(err)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		t := time.NewTicker(60 * time.Second)
		defer t.Stop()

		params := &twitter.HomeTimelineParams{
			Count:          cnt + 1,
			ExcludeReplies: isExcludeReplies,
		}

		for {
			tweets, err := twitterClient.TimelineService.GetHomeTimeline(params)
			if err != nil {
				exitWithError(err)
			}

			for i := len(tweets) - 1; i >= 0; i-- {
				t := tweets[i]

				createdTime, err := t.CreatedAtTime()
				if err != nil {
					exitWithError(err)
				}
				fmt.Printf("%s(@%s): %s\n%s\n\n--------------\n\n", t.User.Name, t.User.ScreenName, createdTime.Format("2006/01/02 15:04:05"), t.Text)
			}

			if !isFollow {
				return
			}

			select {
			case <-t.C:
				if len(tweets) > 0 {
					params.SinceID = tweets[0].ID
				} else {
					fmt.Println("tweets not found")
				}
				continue
			case <-c: // catch signal
				return
			}
		}
	},
}

func init() {
	timelineHomeCmd.Flags().IntP("count", "c", 20, "tweet count")
	timelineHomeCmd.Flags().Bool("exclude-replies", false, "exclude replies")
	timelineCmd.AddCommand(timelineHomeCmd)

	timelineCmd.PersistentFlags().BoolP("follow", "f", false, "Specify if the timelines should be streamed.")
	rootCmd.AddCommand(timelineCmd)
}
