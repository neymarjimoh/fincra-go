package fincra

import (
	"context"
	"errors"
)

const (
	beneficiariesUrl = "/profile/beneficiaries/business/"
)

type BeneficiaryType string
type PaymentDestinationType string

const (
	Individual        BeneficiaryType        = "individual"
	Corporate         BeneficiaryType        = "corporate"
	MobileMoneyWallet PaymentDestinationType = "mobile_money_wallet"
	BankAccount       PaymentDestinationType = "bank_account"
	CryptoWallet      PaymentDestinationType = "crypto_wallet"
	FliqPayWallet     PaymentDestinationType = "fliqpay_wallet"
)

//	beneficiary := CreateBeneficiaryBody{
//		FirstName:         "John",
//		AccountHolderName: "John Doe",
//		Type:              Individual,
//		... other field assignments
//	}
type CreateBeneficiaryBody struct {
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName,omitempty"`
	Email              string `json:"email,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	AccountHolderName  string `json:"accountHolderName"`
	Bank               `json:"bank,omitempty"`
	Type               BeneficiaryType        `json:"type"` // individual or corporate
	Currency           string                 `json:"currency"`
	PaymentDestination PaymentDestinationType `json:"paymentDestination"`
	DestinationAddress string                 `json:"destinationAddress"`
	UniqueIdentifier   string                 `json:"uniqueIdentifier,omitempty"`
	BusinessId         string                 `json:"businessId"` // needed to be passed in params
}

type UpdateBeneficiaryBody struct {
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName,omitempty"`
	Email              string `json:"email,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	AccountHolderName  string `json:"accountHolderName"`
	Bank               `json:"bank,omitempty"`
	Type               BeneficiaryType        `json:"type"` // individual or corporate
	Currency           string                 `json:"currency"`
	PaymentDestination PaymentDestinationType `json:"paymentDestination"`
	DestinationAddress string                 `json:"destinationAddress"`
	UniqueIdentifier   string                 `json:"uniqueIdentifier,omitempty"`
	BusinessId         string                 `json:"businessId"`    // needed to be passed in params
	BeneficiaryId      string                 `json:"beneficiaryId"` // needed to be passed in params
}

type Bank struct {
	Name      string `json:"name,omitempty"`
	Code      string `json:"code,omitempty"`
	SortCode  string `json:"sortCode,omitempty"`
	SwiftCode string `json:"swiftCode,omitempty"`
	Address   `json:"address,omitempty"`
}

type Address struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	Street  string `json:"street,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
}

type createBeneficiaryRequest struct {
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName,omitempty"`
	Email              string `json:"email,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	AccountHolderName  string `json:"accountHolderName"`
	Bank               `json:"bank,omitempty"`
	Type               BeneficiaryType        `json:"type"` // individual or corporate
	Currency           string                 `json:"currency"`
	PaymentDestination PaymentDestinationType `json:"paymentDestination"`
	DestinationAddress string                 `json:"destinationAddress"`
	UniqueIdentifier   string                 `json:"uniqueIdentifier,omitempty"`
}

type GetAllBeneficiariesParams struct {
	BusinessId string `json:"businessId"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
}

type getBeneficiariesRequest struct {
	Page    string `json:"page,omitempty"`
	PerPage string `json:"perPage,omitempty"`
}

type GetBeneficiaryParams struct {
	BusinessId    string `json:"businessId"`
	BeneficiaryId string `json:"beneficiaryId"`
}

// create a beneficiary for business
// client := fincra.NewClient(apiKey)
// resp, err := client.CreateBeneficiary(&client.CreateBeneficiaryBody{})
func (c *Client) CreateBeneficiary(ctx context.Context, beneficiary *CreateBeneficiaryBody) (Response, error) {
	if beneficiary.BusinessId == "" {
		return Response{}, errors.New("business ID is required for beneficiary")
	}

	path := beneficiariesUrl + beneficiary.BusinessId

	// Create a new request object without the BusinessId field
	request := createBeneficiaryRequest{
		FirstName:          beneficiary.FirstName,
		LastName:           beneficiary.LastName,
		Email:              beneficiary.Email,
		PhoneNumber:        beneficiary.PhoneNumber,
		AccountHolderName:  beneficiary.AccountHolderName,
		Bank:               beneficiary.Bank,
		Type:               beneficiary.Type,
		Currency:           beneficiary.Currency,
		PaymentDestination: beneficiary.PaymentDestination,
		DestinationAddress: beneficiary.DestinationAddress,
		UniqueIdentifier:   beneficiary.UniqueIdentifier,
	}

	return c.sendRequest(ctx, "POST", path, &request)
}

// Get all beneficiaries for a business
const businessIdRequiredError = "businessId is required to fetch the beneficiary"

func (c *Client) GetAllBeneficiaries(ctx context.Context, params *GetAllBeneficiariesParams) (Response, error) {
	if params.BusinessId == "" {
		return Response{}, errors.New(businessIdRequiredError)
	}

	if params.Page == "" {
		params.Page = "1"
	}

	if params.PerPage == "" {
		params.PerPage = "10"
	}

	path := beneficiariesUrl + params.BusinessId

	request := getBeneficiariesRequest{
		Page:    params.Page,
		PerPage: params.PerPage,
	}

	return c.sendRequest(ctx, "GET", path, &request)
}

// Get a benefiiciary from a business
func (c *Client) GetBeneficiary(ctx context.Context, params *GetBeneficiaryParams) (Response, error) {
	if params.BusinessId == "" {
		return Response{}, errors.New(businessIdRequiredError)
	}

	if params.BeneficiaryId == "" {
		return Response{}, errors.New("beneficiaryId is required to fetch the beneficiary")
	}

	path := beneficiariesUrl + params.BusinessId + "/" + params.BeneficiaryId

	return c.sendRequest(ctx, "GET", path, nil)
}

// update a beneficiary of a business
// see https://docs.fincra.com/reference/update-a-beneficiary for required parameters
func (c *Client) UpdateBeneficiary(ctx context.Context, body *UpdateBeneficiaryBody) (Response, error) {
	if body.BusinessId == "" {
		return Response{}, errors.New("businessId is required to update the beneficiary")
	}

	if body.BeneficiaryId == "" {
		return Response{}, errors.New("beneficiaryId is required to update the beneficiary")
	}

	path := beneficiariesUrl + body.BusinessId + "/" + body.BeneficiaryId

	request := createBeneficiaryRequest{
		FirstName:          body.FirstName,
		LastName:           body.LastName,
		Email:              body.Email,
		PhoneNumber:        body.PhoneNumber,
		AccountHolderName:  body.AccountHolderName,
		Bank:               body.Bank,
		Type:               body.Type,
		Currency:           body.Currency,
		PaymentDestination: body.PaymentDestination,
		DestinationAddress: body.DestinationAddress,
		UniqueIdentifier:   body.UniqueIdentifier,
	}

	return c.sendRequest(ctx, "PATCH", path, &request)
}

// delete a beneficiary of a business
func (c *Client) DeleteBeneficiary(ctx context.Context, params *GetBeneficiaryParams) (Response, error) {
	if params.BusinessId == "" {
		return Response{}, errors.New("businessId is required to fetch the beneficiary")
	}

	if params.BeneficiaryId == "" {
		return Response{}, errors.New("beneficiaryId is required to fetch the beneficiary")
	}

	path := beneficiariesUrl + params.BusinessId + "/" + params.BeneficiaryId

	return c.sendRequest(ctx, "DELETE", path, nil)
}
