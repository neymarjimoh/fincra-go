package fincra

import "context"

type Type string

const (
	verificationUrl      = "/core/"
	Nuban           Type = "nuban"
	Iban            Type = "iban"
)

type VerifyBankAccountBody struct {
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"bankCode"`
	Type          Type   `json:"type"`
	Iban          string `json:"iban"`
}

type VerifyBVNBody struct {
	Bvn      string `json:"bvn"`
	Business string `json:"business"`
}

// lets you verify a bank account
func (c *Client) VerifyBankAccount(ctx context.Context, verifyData VerifyBankAccountBody) (Response, error) {
	path := verificationUrl + "accounts/resolve"

	return c.sendRequest(ctx, "POST", path, verifyData)
}

// lets you verify a BVN
func (c *Client) VerifyBVN(ctx context.Context, verifyData VerifyBVNBody) (Response, error) {
	path := verificationUrl + "bvn-verification"

	return c.sendRequest(ctx, "POST", path, verifyData)
}
