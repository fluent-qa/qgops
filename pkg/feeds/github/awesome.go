package github

import (
	http "github.com/fluent-qa/qfluent-ops/internal/http"
	"os"
	"strings"
)

// GetGithubAwesomeContent https://raw.githubusercontent.com/luban-agi/Awesome-AIGC-Tutorials/main/README.md
func GetGithubAwesomeContent(githubUrl string) string {
	url := strings.Replace(githubUrl, "github.com", "raw.githubusercontent.com", 1)
	url = strings.Replace(url, ".git", "/main/README.md", 1)
	_, content, _ := http.DefaultClient.Get(url)
	return content
}

func CollectGithubAwesomeContent(category, url string) {
	result := GetGithubAwesomeContent(url)
	_ = os.WriteFile(category+".md", []byte(result), os.ModePerm)
}
