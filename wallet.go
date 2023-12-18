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

// ListWallets lists all the account balance information of a business.
// It takes a context and a business ID as parameters.
// It returns a Response and an error.
func (c *Client) ListWallets(ctx context.Context, businessId string) (Response, error) {

	path := walletUrl + "?businessID=" + businessId

	return c.sendRequest(ctx, "GET", path, nil)
}

// ListWallet provides information to the merchant about a specific account balance.
// It takes a context and an ID as parameters.
// It returns a Response and an error.
func (c *Client) ListWallet(ctx context.Context, id string) (Response, error) {

	path := walletUrl + "/" + id

	return c.sendRequest(ctx, "GET", path, nil)
}

// ListWalletLogs fetches all pay-ins and pay-outs that occurred on your integration.
// It takes a context and a LogsDto as parameters.
// It returns a Response and an error.
func (c *Client) ListWalletLogs(ctx context.Context, logData LogsDto) (Response, error) {

	path := walletUrl + "/logs"

	return c.sendRequest(ctx, "GET", path, logData)
}
