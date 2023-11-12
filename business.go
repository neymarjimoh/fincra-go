package fincra

// fetch the business id and other information of the merchant
// client := fincra.NewClient(apiKey)
// resp, err := client.GetBusinessId()
func (c *Client) GetBusinessId() (Response, error) {
	path := "/profile/business/me"

	return c.sendRequest("GET", path, nil)
}
