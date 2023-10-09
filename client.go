package zenrows

import (
	"net/http"
)

// HttpClient Http client able to perform request, can be http.Client or any other
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client ZenRow client
type Client struct {
	client HttpClient
	config *ClientConfig
}

// NewClient Initialise the client with given HttpClient interface
func NewClient(httpClient HttpClient) *Client {
	config := DefaultConfig()
	return &Client{
		client: httpClient,
		config: &config,
	}
}

// WithApiKey Configures the apikey
func (c *Client) WithApiKey(key string) *Client {
	c.config.ConfigCredentials(key)
	return c
}
