package fincra

import "context"

type TransactionType string
type FeeBearerType string

const (
	Disbursement TransactionType = "disbursement"
	Conversion   TransactionType = "conversion"
	Customer     FeeBearerType   = "customer"
	Business     FeeBearerType   = "business"
)

// CreateQuoteBody represents the request body for creating a quote.
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

// CreateQuote creates a new quote.
// It returns the API response and an error, if any.
func (c *Client) CreateQuote(ctx context.Context, quote *CreateQuoteBody) (Response, error) {
	path := "/quotes/generate"

	return c.sendRequest(ctx, "POST", path, quote)
}
