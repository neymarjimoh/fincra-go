package fincra

import (
	"context"
	"fmt"
	"strings"

	"github.com/neymarjimoh/fincra-go/internal"
)

const (
	virtualAccountUrl = "/profile/virtual-accounts/"
)

// CreateVirtualAccountDto represents the data required to create a virtual account.
type CreateVirtualAccountDto struct {
	Currency               string            `json:"currency"`               // Currency of the virtual account.
	AccountType            string            `json:"accountType"`            // Type of the virtual account.
	KYCInformation         KYCInformationDto `json:"KYCInformation"`         // KYC information of the account holder.
	MeansOfId              []string          `json:"meansOfId"`              // Means of identification of the account holder.
	UtilityBill            string            `json:"utilityBill"`            // Utility bill of the account holder.
	PaymentFlowDescription string            `json:"paymentFlowDescription"` // Description of the payment flow.
	MonthlyVolume          string            `json:"monthlyVolume"`          // Monthly volume of the account.
	EntityName             string            `json:"entityName"`             // Name of the entity associated with the account.
	Attachments            string            `json:"attachments"`            // Attachments related to the account.
	Reason                 string            `json:"reason"`                 // Reason for creating the account.
	Channel                string            `json:"channel"`                // Channel through which the account is created.
}

// KYCInformationDto represents the KYC information of the account holder.
type KYCInformationDto struct {
	FirstName        string      `json:"firstName,omitempty"`        // First name of the account holder.
	LastName         string      `json:"lastName,omitempty"`         // Last name of the account holder.
	Email            string      `json:"email,omitempty"`            // Email of the account holder.
	BusinessName     string      `json:"businessName,omitempty"`     // Business name of the account holder.
	Bvn              string      `json:"bvn,omitempty"`              // BVN (Bank Verification Number) of the account holder.
	BirthDate        string      `json:"birthDate,omitempty"`        // Birth date of the account holder.
	Occupation       string      `json:"occupation,omitempty"`       // Occupation of the account holder.
	BusinessCategory string      `json:"businessCategory,omitempty"` // Business category of the account holder.
	AdditionalInfo   string      `json:"additionalInfo,omitempty"`   // Additional information about the account holder.
	Address          Address     `json:"address,omitempty"`          // Address of the account holder.
	Document         DocumentDto `json:"document,omitempty"`         // Document details of the account holder.
}

// DocumentDto represents the document details of the account holder.
type DocumentDto struct {
	Type              string `json:"type,omitempty"`              // Type of the document.
	Number            string `json:"number,omitempty"`            // Number of the document.
	IssuedCountryCode string `json:"issuedCountryCode,omitempty"` // Country code of the document issuer.
	IssuedBy          string `json:"issuedBy,omitempty"`          // Issuer of the document.
	IssuedDate        string `json:"issuedDate,omitempty"`        // Date of issuance of the document.
	ExpirationDate    string `json:"expirationDate,omitempty"`    // Expiration date of the document.
}

// Options represents the filtering options for listing virtual accounts.
type Options struct {
	Currency      string `json:"currency"`      // Currency of the virtual accounts.
	BusinessName  string `json:"businessName"`  // Business name associated with the virtual accounts.
	IssuedDate    string `json:"issuedDate"`    // Issued date of the virtual accounts.
	RequestedDate string `json:"requestedDate"` // Requested date of the virtual accounts.
	AccountNumber string `json:"accountNumber"` // Account number of the virtual accounts.
	Status        string `json:"status"`        // Status of the virtual accounts.
}

// CreateVirtualAccount creates a new virtual account.
// Read more: [https://docs.fincra.com/reference/request-virtual-accounts]
func (c *Client) CreateVirtualAccount(ctx context.Context, data CreateVirtualAccountDto) (Response, error) {
	path := virtualAccountUrl + "/requests"
	return c.sendRequest(ctx, "POST", path, data)
}

// ListVirtualAccounts lists virtual accounts based on the provided options.
// At least one option (currency, businessName, issuedDate, requestedDate, accountNumber, status) must be specified.
// Read more: [https://docs.fincra.com/reference/get-merchant-virtual-account-requests]
func (c *Client) ListVirtualAccounts(ctx context.Context, options Options) (Response, error) {
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

	return c.sendRequest(ctx, "GET", path, nil)
}

// ListVirtualAccountRequests lists all account requests belonging to a merchant.
func (c *Client) ListVirtualAccountRequests(ctx context.Context) (Response, error) {
	path := virtualAccountUrl + "requests"
	return c.sendRequest(ctx, "GET", path, nil)
}

// ListVirtualAccountByCurrency retrieves virtual accounts belonging to a merchant by currency.
func (c *Client) ListVirtualAccountByCurrency(ctx context.Context, currency string) (Response, error) {
	path := virtualAccountUrl + "?currency=" + currency
	return c.sendRequest(ctx, "GET", path, nil)
}

// ListVirtualAccountByBvn retrieves virtual accounts belonging to a merchant by BVN (Bank Verification Number).
func (c *Client) ListVirtualAccountByBvn(ctx context.Context, bvn, businessId string) (Response, error) {
	path := virtualAccountUrl + "/bvn?bvn=" + bvn + "&business=" + businessId
	return c.sendRequest(ctx, "GET", path, nil)
}

// ListVirtualAccount retrieves a virtual account by its account ID.
func (c *Client) ListVirtualAccount(ctx context.Context, accountId string) (Response, error) {
	path := virtualAccountUrl + "/virtual-accounts/" + accountId
	return c.sendRequest(ctx, "GET", path, nil)
}
