package feishu

import (
	"fmt"
	feishu2 "github.com/fluent-qa/qfluent-ops/pkg/tableapp/feishu"
	"os"
)

type ConfigOpts struct {
	appId     string
	appSecret string
}

var configOpts = ConfigOpts{}

func handleConfigCommand(opts *ConfigOpts) error {
	configPath, err := feishu2.GetConfigFilePath()
	feishu2.CheckErr(err)

	fmt.Println("Configuration file on: " + configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := feishu2.NewConfig(opts.appId, opts.appSecret)
		if err = config.WriteConfig2File(configPath); err != nil {
			return err
		}
		fmt.Println(feishu2.PrettyPrint(config))
	} else {
		config, err := feishu2.ReadConfigFromFile(configPath)
		if err != nil {
			return err
		}
		if opts.appId != "" {
			config.Feishu.AppId = opts.appId
		}
		if opts.appSecret != "" {
			config.Feishu.AppSecret = opts.appSecret
		}
		if opts.appId != "" || opts.appSecret != "" {
			if err = config.WriteConfig2File(configPath); err != nil {
				return err
			}
		}
		fmt.Println(feishu2.PrettyPrint(config))
	}
	return nil
}
