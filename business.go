package fincra

import "context"

// GetBusinessId fetches the business id and other information of the merchant
// client := fincra.NewClient(apiKey)
// resp, err := client.GetBusinessId()
func (c *Client) GetBusinessId(ctx context.Context) (Response, error) {
	path := "/profile/business/me"

	return c.sendRequest(ctx, "GET", path, nil)
}
