package config

type DiscordConf struct {
	DISCORD_USER_TOKEN string
	DISCORD_BOT_TOKEN  string
	DISCORD_SERVER_ID  string
	DISCORD_CHANNEL_ID string
	CB_URL             string
}

type OpenAIConfig struct {
	Keys     map[string][]string `json:"keys"`
	Helicone string              `json:"helicone"`
	BaseURL  string              `json:"baseurl"`
}
