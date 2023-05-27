package fincra

import (
	"encoding/json"
	"errors"
)

type BeneficiaryType string
type PaymentDestinationType string

const (
	Individual        BeneficiaryType        = "individual"
	Corporate         BeneficiaryType        = "corporate"
	MobileMoneyWallet PaymentDestinationType = "mobile_money_wallet"
	BankAccount       PaymentDestinationType = "bank_account"
	CryptoWallet      PaymentDestinationType = "crypto_wallet"
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
	BusinessId string `json:"businessId"`
	BeneficiaryId       string `json:"beneficiaryId"`
}

// create a beneficiary for business
// client := fincra.NewClient(apiKey)
// resp, err := client.CreateBeneficiary(&client.CreateBeneficiaryBody{})
func (c *Client) CreateBeneficiary(beneficiary *CreateBeneficiaryBody) (Response, error) {
	if beneficiary.BusinessId == "" {
		return Response{}, errors.New("business ID is required for beneficiary")
	}

	path := "/profile/beneficiaries/business/" + beneficiary.BusinessId

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

	response, err := c.sendRequest("POST", path, &request)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

func (c *Client) GetAllBeneficiaries(params *GetAllBeneficiariesParams) (Response, error) {
	if params.BusinessId == "" {
		return Response{}, errors.New("businessId is required to fetch the beneficiary")
	}

	if params.Page == "" {
		params.Page = "1"
	}

	if params.PerPage == "" {
		params.PerPage = "10"
	}

	path := "/profile/beneficiaries/business/" + params.BusinessId

	request := getBeneficiariesRequest{
		Page:    params.Page,
		PerPage: params.PerPage,
	}

	response, err := c.sendRequest("GET", path, &request)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

func (c *Client) GetBeneficiary(params *GetBeneficiaryParams) (Response, error) {
	if params.BusinessId == "" {
		return Response{}, errors.New("businessId is required to fetch the beneficiary")
	}

	if params.BeneficiaryId == "" {
		return Response{}, errors.New("beneficiaryId is required to fetch the beneficiary")
	}

	path := "/profile/beneficiaries/business/" + params.BusinessId + "/" + params.BeneficiaryId

	response, err := c.sendRequest("GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
