package http

import (
	"fmt"
	"testing"
)

const MD = "https://raw.githubusercontent.com/luban-agi/Awesome-AIGC-Tutorials/main/README.md"
const MD_ZH = "https://raw.githubusercontent.com/luban-agi/Awesome-AIGC-Tutorials/main/README_zh.md"

var client = NewHttpClient()

func TestHttpClient_Get(t *testing.T) {
	statusCode, md, _ := client.Get(MD)
	statusCode1, md_zh, _ := client.Get(MD_ZH)
	fmt.Println(statusCode, md)
	fmt.Println(statusCode1, md_zh)
}
