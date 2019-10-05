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

func LoadConfigFromViper() (*Config, error) {
	var config *Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
