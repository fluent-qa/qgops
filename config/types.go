package config

type OpenAIConfig struct {
	Keys     map[string][]string `json:"keys"`
	Helicone string              `json:"helicone"`
	BaseURL  string              `json:"baseurl"`
}
