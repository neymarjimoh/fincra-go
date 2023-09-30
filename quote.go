package fincra

import (
	"encoding/json"
)

type TransactionType string
type FeeBearerType string

const (
	Disbursement TransactionType = "disbursement"
	Conversion   TransactionType = "conversion"
	Customer     FeeBearerType   = "customer"
	Business     FeeBearerType   = "business"
)

type CreateQuoteBody struct {
	Action              string                 `json:"action"`
	TransactionType     TransactionType        `json:"transactionType"`
	FeeBearer           FeeBearerType          `json:"feeBearer"`
	PaymentDestination  PaymentDestinationType `json:"paymentDestination"`
	BeneficiaryType     BeneficiaryType        `json:"beneficiaryType"`
	PaymentScheme       string                 `json:"paymentScheme,omitempty"`
	Business            string                 `json:"business"`
	Amount              string                 `json:"amount"`
	DestinationCurrency string                 `json:"destinationCurrency"`
	SourceCurrency      string                 `json:"sourceCurrency"`
}

func (c *Client) CreateQuote(quote *CreateQuoteBody) (Response, error) {
	path := "/quotes/generate"

	response, err := c.sendRequest("POST", path, quote)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
