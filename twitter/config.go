package twitter

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConsumerApiKey    string
	ConsumerApiSecret string
	AccessToken       string
	AccessSecret      string
}

func LoadConfigFromViper() *Config {
	return &Config{
		ConsumerApiKey:    viper.GetString("consumer_api_key"),
		ConsumerApiSecret: viper.GetString("consumer_api_secret"),
		AccessToken:       viper.GetString("access_token"),
		AccessSecret:      viper.GetString("access_secret"),
	}
}
