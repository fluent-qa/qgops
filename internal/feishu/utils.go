package feishu

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func CheckErr(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		fmt.Fprintf(
			os.Stderr,
			"\n%s\n\n%s\n\n",
			strings.Repeat("=", 20),
			"Report the following if it is a bug",
		)
		panic(e)
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func UnescapeURL(rawURL string) string {
	if u, err := url.QueryUnescape(rawURL); err == nil {
		return u
	}
	return rawURL
}

func ValidateDownloadURL(url, allowHost string) (string, string, string, error) {
	hosts := []string{"feishu.cn", "larksuite.com"}
	if allowHost != "" {
		hosts = append(hosts, allowHost)
	}

	reg := regexp.MustCompile("^https://([\\w-]+.)?(" + strings.Join(hosts, "|") + ")/(docs|docx|wiki)/([a-zA-Z0-9]+)")
	matchResult := reg.FindStringSubmatch(url)
	if matchResult == nil || len(matchResult) != 5 {
		return "", "", "", errors.Errorf("Invalid feishu/larksuite/allowHost URL format")
	}
	return matchResult[2], matchResult[3], matchResult[4], nil
}

const projectDirName = "feishu2md"

func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal("Can not load .env file")
		os.Exit(-1)
	}
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "..")
	return root
}
