package qhttp

//https://github.com/imroc/req

import (
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"testing"
	"time"
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

func TestReqMode(t *testing.T) {
	req.DevMode()
	req.EnableAutoDecode()
	req.EnableForceHTTP1()
	var result = req.MustGet("https://httpbin.org/uuid")
	print(result)
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (msg *ErrorMessage) Error() string {
	return fmt.Sprintf("API Error: %s", msg.Message)
}

type UserInfo struct {
	Name string `json:"name"`
	Blog string `json:"blog"`
}

func TestAdvancedUsage(t *testing.T) {
	client := req.C().
		SetUserAgent("my-custom-client"). // Chainable client settings.
		SetTimeout(5 * time.Second).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil { // There is an underlying error, e.g. network error or unmarshal error.
				return nil
			}
			if errMsg, ok := resp.ErrorResult().(*ErrorMessage); ok {
				resp.Err = errMsg // Convert api error into go error
				return nil
			}
			if !resp.IsSuccessState() {
				// Neither a success response nor a error response, record details to help troubleshooting
				resp.Err = fmt.Errorf("bad status: %s\nraw content:\n%s", resp.Status, resp.Dump())
			}
			return nil
		})

	var userInfo UserInfo
	var errMsg ErrorMessage
	resp, err := client.R().
		SetHeader("Accept", "application/vnd.github.v3+json"). // Chainable request settings.
		SetPathParam("username", "imroc").                     // Replace path variable in url.
		SetSuccessResult(&userInfo).                           // Unmarshal response body into userInfo automatically if status code is between 200 and 299.
		SetErrorResult(&errMsg).                               // Unmarshal response body into errMsg automatically if status code >= 400.
		EnableDump().                                          // Enable dump at request level, only print dump content if there is an error or some unknown situation occurs to help troubleshoot.
		Get("https://api.github.com/users/{username}")

	if err != nil { // Error handling.
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump()) // Record raw content when error occurs.
		return
	}

	if resp.IsErrorState() { // Status code >= 400.
		fmt.Println(errMsg.Message) // Record error message returned.
		return
	}

	if resp.IsSuccessState() { // Status code is between 200 and 299.
		fmt.Printf("%s (%s)\n", userInfo.Name, userInfo.Blog)
		return
	}

	// Unknown status code.
	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump()) // Record raw content when server returned unknown status code.
}

type Repo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Result struct {
	Data string `json:"data"`
}

func TestSimplePost(t *testing.T) {
	client := req.C().DevMode()
	var result Result

	resp, err := client.R().
		SetBody(&Repo{Name: "req", Url: "https://github.com/imroc/req"}).
		SetSuccessResult(&result).
		Post("https://httpbin.org/post")
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
		return
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("data:", result.Data)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
}

type APIResponse struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func TestApiStyleReq(t *testing.T) {
	var resp APIResponse
	c := req.C().SetBaseURL("https://httpbin.org/post")
	err := c.Post().
		SetBody("hello").
		Do().
		Into(&resp)
	if err != nil {
		panic(err)
	}
	fmt.Println("My IP is", resp.Origin)
}
