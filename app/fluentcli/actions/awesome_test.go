package actions

import (
	"fmt"
	"github.com/samber/lo"
	"log/slog"
	"testing"
)

func Test_readConfigFile(t *testing.T) {

	got, err := readConfigFile("awesome.json")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(got)
}

func Test_writeToLocal(t *testing.T) {
	got, err := readConfigFile("awesome.json")
	if err != nil {
		t.Error(err)
	}
	lo.ForEach(got.Awesomes, func(item AwesomeRepo, index int) {
		err := writeToLocalFolder("awesome", item)
		if err != nil {
			slog.Error(err.Error())
		}
	})
}

func Test_AddUrl(t *testing.T) {
	AddAwesomeURL("ai", "https://github.com/hyp1231/awesome-llm-powered-agent.git")
}

func Test_fetch_All(t *testing.T) {
	FetchAllAwesome()
}
