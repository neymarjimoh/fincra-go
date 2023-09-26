package fincra

import (
	"context"
	"encoding/json"
)

// fetch the business id and other information of the merchant
// client := fincra.NewClient(apiKey)
// resp, err := client.GetBusinessId()
func (c *Client) GetBusinessId() (Response, error) {
	path := "/profile/business/me"

	ctx, cancel := context.WithTimeout(context.Background(), c.HttpClient.Timeout)
	defer cancel()

	response, err := c.sendRequest(ctx, "GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
