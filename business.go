package fincra

import "encoding/json"

// fetch the business id and other information of the merchant
// client := fincra.NewClient(apiKey)
// resp, err := client.GetBusinessId()
func (c *Client) GetBusinessId() (Response, error) {
	path := "/profile/business/me"

	response, err := c.sendRequest("GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
