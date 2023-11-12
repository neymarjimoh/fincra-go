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

type Client struct {
	apiKey     string
	BaseUrl    *url.URL
	HttpClient *http.Client
}

type Option func(c *Client)

func WithSandbox(isSandbox bool) Option {
	return func(c *Client) {
		c.BaseUrl = getBaseUrl(isSandbox)
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.HttpClient.Timeout = timeout
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

type Response map[string]interface{}

func (c *Client) sendRequest(method, path string, payload interface{}, ctx ...context.Context) (Response, error) {
	defaultCtx, cancel := context.WithTimeout(context.Background(), c.HttpClient.Timeout)
	defer cancel()

	var effectiveCtx context.Context

	if len(ctx) > 0 {
		effectiveCtx = ctx[0]
	} else {
		effectiveCtx = defaultCtx
	}

	var buf io.ReadWriter
	if payload != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	}

	url := fmt.Sprintf("%s%s", c.BaseUrl.String(), path)

	req, err := http.NewRequestWithContext(effectiveCtx, method, url, buf)
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
	_ = json.Unmarshal(responseBody, &jsonResponse)

	return jsonResponse, nil
}
