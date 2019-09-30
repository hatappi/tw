package twitter

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConsumerApiKey    string `mapstructure:"consumer_api_key"`
	ConsumerApiSecret string `mapstructure:"consumer_api_secret"`
	AccessToken       string `mapstructure:"access_token"`
	AccessSecret      string `mapstructure:"access_secret"`
}

func LoadConfigFromViper() *Config {
	var config *Config

	viper.Unmarshal(&config)

	return config
}
