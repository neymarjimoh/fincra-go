package fincra

import (
	"encoding/json"
)

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
func (c *Client) VerifyBankAccount(verifyData VerifyBankAccountBody) (Response, error) {
	path := verificationUrl + "accounts/resolve"

	response, err := c.sendRequest("POST", path, verifyData)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

// lets you verify a BVN
func (c *Client) VerifyBVN(verifyData VerifyBVNBody) (Response, error) {
	path := verificationUrl + "bvn-verification"

	response, err := c.sendRequest("POST", path, verifyData)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
