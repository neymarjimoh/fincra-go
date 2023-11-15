package fincra

import (
	"context"
	"errors"
)

const converisionsUrl = "/conversions"

type CreateConversionBody struct {
	BusinessId     string `json:"business"`
	QuoteReference string `json:"quoteReference"`
}

func (c *Client) GetBusinessConversions(ctx context.Context, businessId string) (Response, error) {
	path := converisionsUrl + "?business=" + businessId

	return c.sendRequest(ctx, "GET", path, nil)
}

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

func (c *Client) GetConversion(ctx context.Context, conversionId string) (Response, error) {
	path := converisionsUrl + "/" + conversionId

	return c.sendRequest(ctx, "GET", path, nil)
}
