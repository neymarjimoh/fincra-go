package fincra

import "encoding/json"

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

type KYCInformation struct {
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

// read here https://docs.fincra.com/reference/request-virtual-accounts before using this method
func (c *Client) CreateVirtualAccount(data CreateVirtualAccountDto) (Response, error) {

	path := virtualAccountUrl + "/requests"

	response, err := c.sendRequest("POST", path, data)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
