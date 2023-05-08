package fincra

import (
	"net/http"
	"net/url"
	"time"
)

const (
	sandboxUrl = "https://sandboxapi.fincra.com"
	liveUrl = "https://api.fincra.com"
)

type Client struct {
	apiKey string
	BaseUrl *url.URL
	HttpClient *http.Client
}

type Option func(c *Client)

func WithSandbox(isSandbox bool) Option {
	return func(c *Client) {
		c.BaseUrl = getBaseUrl(isSandbox)
	}
}

// instantiate a new fincra client
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey: apiKey,
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
		BaseUrl: getBaseUrl(false),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// gets the correct base url for live and test mode
func getBaseUrl(sandbox bool) *url.URL {
	url, _ := url.Parse(liveUrl)
	if sandbox {
		url, _ = url.Parse(sandboxUrl)
		return url
	}
	return url
}
