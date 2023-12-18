package fincra

import (
	"context"
	"errors"
)

const converisionsUrl = "/conversions"

// CreateConversionBody represents the request body for creating a conversion.
type CreateConversionBody struct {
	BusinessId     string `json:"business"`
	QuoteReference string `json:"quoteReference"`
}

// GetBusinessConversions retrieves the conversions associated with a specific business.
// It takes a context and a businessId as parameters.
// It returns a Response and an error.
func (c *Client) GetBusinessConversions(ctx context.Context, businessId string) (Response, error) {
	path := converisionsUrl + "?business=" + businessId

	return c.sendRequest(ctx, "GET", path, nil)
}

// CreateConversion initiates a new conversion.
// It takes a context and a CreateConversionBody as parameters.
// It returns a Response and an error.
// If the businessId or quoteReference is empty, it returns an error.
func (c *Client) CreateConversion(ctx context.Context, conversion *CreateConversionBody) (Response, error) {
	if conversion.BusinessId == "" {
		return Response{}, errors.New("business is required to convert a currency")
	}

	if conversion.QuoteReference == "" {
		return Response{}, errors.New("quoteReference is required to convert a currency")
	}
	path := converisionsUrl + "/initiate"

	return c.sendRequest(ctx, "POST", path, conversion)
}

// GetConversion retrieves a specific conversion by its ID.
// It takes a context and a conversionId as parameters.
// It returns a Response and an error.
func (c *Client) GetConversion(ctx context.Context, conversionId string) (Response, error) {
	path := converisionsUrl + "/" + conversionId

	return c.sendRequest(ctx, "GET", path, nil)
}
