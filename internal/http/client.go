package http

import (
	"github.com/imroc/req/v3"
)

type Client struct {
	Client  *req.Client
	Options map[string]string
}

var DefaultClient = NewHttpClient()

func NewHttpClient() *Client {
	return &Client{
		Client:  req.DefaultClient(),
		Options: nil,
	}
}

func (h Client) Get(url string) (string, string, *req.Response) {
	resp := h.Client.R(). // Use R() to create a request.
				MustGet(url)
	return resp.Status, resp.String(), resp
}
