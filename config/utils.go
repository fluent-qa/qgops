package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func getViperStringValue(key string) string {
	value := viper.GetString(key)
	if value == "" {
		panic(fmt.Errorf("%s MUST be provided in environment or DiscordConfig.yaml file", key))
	}
	return value
}
