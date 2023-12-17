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

// VerifyBankAccount is a method that allows you to verify a bank account.
// It takes a context and a VerifyBankAccountBody as input parameters.
// It returns a Response and an error as output parameters.
func (c *Client) VerifyBankAccount(ctx context.Context, verifyData VerifyBankAccountBody) (Response, error) {
	path := verificationUrl + "accounts/resolve"

	return c.sendRequest(ctx, "POST", path, verifyData)
}

// VerifyBVN is a method that allows you to verify a BVN (Bank Verification Number).
// It takes a context and a VerifyBVNBody as input parameters.
// It returns a Response and an error as output parameters.
func (c *Client) VerifyBVN(ctx context.Context, verifyData VerifyBVNBody) (Response, error) {
	path := verificationUrl + "bvn-verification"

	return c.sendRequest(ctx, "POST", path, verifyData)
}
