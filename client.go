package zenrows

import (
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	client HttpClient
	config *ClientConfig
}

func NewClient(httpClient HttpClient) *Client {
	config := DefaultConfig()
	return &Client{
		client: httpClient,
		config: &config,
	}
}

func (c *Client) WithApiKey(key string) *Client {
	c.config.ConfigCredentials(key)
	return c
}
