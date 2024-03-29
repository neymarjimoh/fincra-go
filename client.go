package fincra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	sandboxUrl = "https://sandboxapi.fincra.com"
	liveUrl    = "https://api.fincra.com"
)

// Client represents a client for interacting with the Fincra API.
type Client struct {
	apiKey     string
	BaseUrl    *url.URL
	HttpClient *http.Client
}

// Option represents an option for configuring the Fincra client.
type Option func(c *Client)

// WithSandbox sets the client to use the sandbox environment.
func WithSandbox(isSandbox bool) Option {
	return func(c *Client) {
		c.BaseUrl = getBaseUrl(isSandbox)
	}
}

// WithTimeout sets the timeout duration for HTTP requests made by the client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.HttpClient.Timeout = timeout
	}
}

// NewClient creates a new instance of the Fincra client.
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

func getBaseUrl(sandbox bool) *url.URL {
	url, _ := url.Parse(liveUrl)
	if sandbox {
		url, _ = url.Parse(sandboxUrl)
		return url
	}
	return url
}

// Response represents the response from the Fincra API.
type Response map[string]interface{}

func (c *Client) sendRequest(ctx context.Context, method, path string, payload interface{}) (Response, error) {
	var buf io.ReadWriter
	if payload != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	}

	url := fmt.Sprintf("%s%s", c.BaseUrl.String(), path)

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("error instantiating request: %w", err)
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.apiKey)

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response from body: %w", err)
	}

	var jsonResponse Response
	if err := json.Unmarshal(responseBody, &jsonResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return jsonResponse, nil
}
