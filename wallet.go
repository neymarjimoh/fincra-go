package fincra

import "context"

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
func (c *Client) ListWallets(ctx context.Context, businessId string) (Response, error) {

	path := walletUrl + "?businessID=" + businessId

	return c.sendRequest(ctx, "GET", path, nil)
}

// provides information to the merchant about a specific account balance
func (c *Client) ListWallet(ctx context.Context, id string) (Response, error) {

	path := walletUrl + "/" + id

	return c.sendRequest(ctx, "GET", path, nil)
}

// fetches all pay-ins and pay-outs that occurred on your integration
func (c *Client) ListWalletLogs(ctx context.Context, logData LogsDto) (Response, error) {

	path := walletUrl + "/logs"

	return c.sendRequest(ctx, "GET", path, logData)
}
