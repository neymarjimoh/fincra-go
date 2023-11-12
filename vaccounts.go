package fincra

import (
	"fmt"
	"strings"

	"github.com/neymarjimoh/fincra-go/utils"
)

const (
	virtualAccountUrl = "/profile/virtual-accounts/"
)

type CreateVirtualAccountDto struct {
	Currency               string            `json:"currency"`
	AccountType            string            `json:"accountType"`
	KYCInformation         KYCInformationDto `json:"KYCInformation"`
	MeansOfId              []string          `json:"meansOfId"`
	UtilityBill            string            `json:"utilityBill"`
	PaymentFlowDescription string            `json:"paymentFlowDescription"`
	MonthlyVolume          string            `json:"monthlyVolume"`
	EntityName             string            `json:"entityName"`
	Attachments            string            `json:"attachments"`
	Reason                 string            `json:"reason"`
	Channel                string            `json:"channel"`
}

type KYCInformationDto struct {
	FirstName        string `json:"firstName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	Email            string `json:"email,omitempty"`
	BusinessName     string `json:"businessName,omitempty"`
	Bvn              string `json:"bvn,omitempty"`
	BirthDate        string `json:"birthDate,omitempty"`
	Occupation       string `json:"occupation,omitempty"`
	BusinessCategory string `json:"businessCategory,omitempty"`
	AdditionalInfo   string `json:"additionalInfo,omitempty"`
	Address          `json:"address,omitempty"`
	Document         DocumentDto `json:"document,omitempty"`
}

type DocumentDto struct {
	Type              string `json:"type,omitempty"`
	Number            string `json:"number,omitempty"`
	IssuedCountryCode string `json:"issuedCountryCode,omitempty"`
	IssuedBy          string `json:"issuedBy,omitempty"`
	IssuedDate        string `json:"issuedDate,omitempty"`
	ExpirationDate    string `json:"expirationDate,omitempty"`
}

type Options struct {
	Currency      string `json:"currency"`
	BusinessName  string `json:"businessName"`
	IssuedDate    string `json:"issuedDate"`
	RequestedDate string `json:"requestedDate"`
	AccountNumber string `json:"accountNumber"`
	Status        string `json:"status"`
}

// read here https://docs.fincra.com/reference/request-virtual-accounts before using this method
func (c *Client) CreateVirtualAccount(data CreateVirtualAccountDto) (Response, error) {

	path := virtualAccountUrl + "/requests"

	return c.sendRequest("POST", path, data)
}

// read here https://docs.fincra.com/reference/get-merchant-virtual-account-requests for more info on this method
func (c *Client) ListVirtualAccounts(options Options) (Response, error) {
	if utils.IsEmpty(options) {
		return nil, fmt.Errorf("at least one option (currency, businessName, issuedDate, requestedDate, accountNumber, status) must be specified")
	}

	var queryParameters []string

	if options.Currency != "" {
		queryParameters = append(queryParameters, "currency="+options.Currency)
	}
	if options.BusinessName != "" {
		queryParameters = append(queryParameters, "businessName="+options.BusinessName)
	}
	if options.IssuedDate != "" {
		queryParameters = append(queryParameters, "issuedDate="+options.IssuedDate)
	}
	if options.RequestedDate != "" {
		queryParameters = append(queryParameters, "requestedDate="+options.RequestedDate)
	}
	if options.AccountNumber != "" {
		queryParameters = append(queryParameters, "accountNumber="+options.AccountNumber)
	}
	if options.Status != "" {
		queryParameters = append(queryParameters, "status="+options.Status)
	}

	path := virtualAccountUrl
	if len(queryParameters) > 0 {
		path += "?" + strings.Join(queryParameters, "&")
	}

	return c.sendRequest("GET", path, nil)
}

// This method is used for getting all account requests belonging to a merchant
func (c *Client) ListVirtualAccountRequests() (Response, error) {
	path := virtualAccountUrl + "requests"

	return c.sendRequest("GET", path, nil)
}

// This method is used for retrieving an account that is belongs to a merchant by currency
func (c *Client) ListVirtualAccountByCurrency(currency string) (Response, error) {
	path := virtualAccountUrl + "?currency=" + currency

	return c.sendRequest("GET", path, nil)
}

// This method is used for retrieving an account that is belongs to a merchant by BVN
func (c *Client) ListVirtualAccountByBvn(bvn, businessId string) (Response, error) {
	path := virtualAccountUrl + "/bvn?bvn=" + bvn + "&business=" + businessId

	return c.sendRequest("GET", path, nil)
}

// This method is used for retrieving a virtual account
func (c *Client) ListVirtualAccount(accountId string) (Response, error) {
	path := virtualAccountUrl + "/virtual-accounts/" + accountId

	return c.sendRequest("GET", path, nil)
}
