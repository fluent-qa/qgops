package actions

import (
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/imroc/req/v3"

	"github.com/fluent-qa/qgops/internal/utils/jsonutil"
	_ "github.com/fluent-qa/qgops/internal/utils/qhttp"
)

var (
	FetchAllCmd = &cobra.Command{
		Use:   "awesome-all",
		Short: "Fetch all awesome repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			return FetchAllAwesome()
		},
	}

	AddURLCmd = &cobra.Command{
		Use:   "awesome-add <category> <github-url>",
		Short: "Add a new awesome repository URL",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return AddAwesomeURL(args[0], args[1])
		},
	}

	FetchByCategoryCmd = &cobra.Command{
		Use:   "awesome <category>",
		Short: "Fetch awesome repositories by category",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			FetchAwesomeByCategory(args[0])
		},
	}
)

// TODO: move to sqlite database instead of a configuration file

type AwesomeRepo struct {
	Category  string `json:"category"`
	GithubURL string `json:"github_url"`
}

type AwesomeRepos struct {
	Awesomes  []AwesomeRepo `json:"awesomes,omitempty"`
	OutputDir string        `json:"outputDir,omitempty"`
}

func FetchAllAwesome() error {
	repos, _ := readConfigFile(AwesomeRepoConfig)
	lo.ForEach(repos.Awesomes, func(repo AwesomeRepo, index int) {
		_ = writeToLocalFolder(repos.OutputDir, repo)
	})
	return nil
}

func AddAwesomeURL(category, githubUrl string) error {
	repos, _ := readConfigFile(AwesomeRepoConfig)
	repos.Awesomes = append(repos.Awesomes, AwesomeRepo{
		Category: category, GithubURL: githubUrl,
	})
	newConfig := jsonutil.ToString(repos)
	_ = os.WriteFile(AwesomeRepoConfig, []byte(newConfig), os.ModePerm.Perm())
	slog.Info("Add new awesome {category},{url}", category, githubUrl)
	return nil
}

func FetchAwesomeByCategory(category string) {
	repos, _ := readConfigFile(AwesomeRepoConfig)
	categoryRepos := lo.Filter(repos.Awesomes, func(item AwesomeRepo, index int) bool {
		if item.Category == category {
			return true
		}
		return false
	})
	lo.ForEach(categoryRepos, func(repo AwesomeRepo, index int) {
		_ = writeToLocalFolder(repos.OutputDir, repo)
	})
}

func readConfigFile(filePath string) (*AwesomeRepos, error) {
	awesomeRepoConfig := &AwesomeRepos{}
	jsonutil.ToStructFromFile(filePath, awesomeRepoConfig)
	return awesomeRepoConfig, nil
}

func writeToLocalFolder(outputDirs string, repo AwesomeRepo) error {
	githubReadMeURL, repoName := convertToRawReadMeURL(repo.GithubURL)
	resp, err := req.C().R().Get(githubReadMeURL)
	//https://raw.githubusercontent.com/hyp1231/awesome-llm-powered-agent/refs/heads/main/README.md
	//https://raw.githubusercontent.com/hyp1231/awesome-llm-powered-agent/refs/heads/master/README.md
	if resp.StatusCode == 404 {
		resp, err = req.C().R().Get(strings.ReplaceAll(githubReadMeURL, "/main/", "/master/"))
	}
	if err != nil {
		return fmt.Errorf("failed to Get ReadMe file: %v", err)
	}
	// Create directory
	dirPath := filepath.Join(outputDirs, repo.Category)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Create file name
	fileName := repoName + ".md"
	readMeFilePath := filepath.Join(dirPath, fileName)
	file, err := os.OpenFile(readMeFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if _, err = file.WriteString(string(resp.String())); err != nil {
		return err
	}
	return nil
}

func convertToRawReadMeURL(githubUrl string) (string, string) {
	// Remove .git from the end if present
	githubUrl = strings.ReplaceAll(githubUrl, ".git", "")
	parts := strings.Split(githubUrl, "/")
	repoName := ""
	if len(parts) >= 2 {
		repoName = parts[len(parts)-1]
	}
	githubUrl = strings.ReplaceAll(githubUrl, " ", "")
	// Replace github.com with raw.githubusercontent.com
	rawURL := strings.Replace(githubUrl, "github.com", "raw.githubusercontent.com", 1)
	// Add /master/README.md to the end
	rawURL += "/refs/heads/main/README.md"

	return rawURL, repoName
}
