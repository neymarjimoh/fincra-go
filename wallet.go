package fincra

import (
	"encoding/json"
)

type ActionType string

const (
	walletUrl            = "/wallets"
	Credit    ActionType = "credit"
	Debit     ActionType = "debit"
)

type LogsDto struct {
	Business string     `json:"business"`
	Action   ActionType `json:"action"`
	Amount   string     `json:"amount"`
	Page     string     `json:"page"`
	PerPage  string     `json:"perPage"`
}

// lists all the account balance information of a business
func (c *Client) ListWallets(businessId string) (Response, error) {

	path := walletUrl + "?businessID=" + businessId

	response, err := c.sendRequest("GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

// provides information to the merchant about a specific account balance
func (c *Client) ListWallet(id string) (Response, error) {

	path := walletUrl + "/" + id

	response, err := c.sendRequest("GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

// fetches all pay-ins and pay-outs that occurred on your integration
func (c *Client) ListWalletLogs(logData LogsDto) (Response, error) {

	path := walletUrl + "/logs"

	response, err := c.sendRequest("GET", path, logData)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
