package config

import (
	"github.com/spf13/viper"
)

type DiscordConf struct {
	DISCORD_USER_TOKEN string
	DISCORD_BOT_TOKEN  string
	DISCORD_SERVER_ID  string
	DISCORD_CHANNEL_ID string
	CB_URL             string
}

var DiscordConfig *DiscordConf

func LoadDiscordConfig(cfg string) (*DiscordConf, error) {
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	DiscordConfig = &DiscordConf{
		DISCORD_USER_TOKEN: getViperStringValue("DISCORD_USER_TOKEN"),
		DISCORD_BOT_TOKEN:  getViperStringValue("DISCORD_BOT_TOKEN"),
		DISCORD_SERVER_ID:  getViperStringValue("DISCORD_SERVER_ID"),
		DISCORD_CHANNEL_ID: getViperStringValue("DISCORD_CHANNEL_ID"),
		CB_URL:             getViperStringValue("CB_URL"),
	}
	return DiscordConfig, nil
}

func GetDiscordConfig() *DiscordConf {
	return DiscordConfig
}
