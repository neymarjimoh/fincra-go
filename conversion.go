package fincra

import (
	"context"
	"encoding/json"
	"errors"
)

const converisionsUrl = "/conversions"

type CreateConversionBody struct {
	BusinessId     string `json:"business"`
	QuoteReference string `json:"quoteReference"`
}

func (c *Client) GetBusinessConversions(businessId string) (Response, error) {
	path := converisionsUrl + "?business=" + businessId

	ctx, cancel := context.WithTimeout(context.Background(), c.HttpClient.Timeout)
	defer cancel()

	response, err := c.sendRequest(ctx, "GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

func (c *Client) CreateConversion(conversion *CreateConversionBody) (Response, error) {
	if conversion.BusinessId == "" {
		return Response{}, errors.New("business is required to convert a currency")
	}

	if conversion.QuoteReference == "" {
		return Response{}, errors.New("quoteReference is required to convert a currency")
	}
	path := converisionsUrl + "/initiate"

	ctx, cancel := context.WithTimeout(context.Background(), c.HttpClient.Timeout)
	defer cancel()

	response, err := c.sendRequest(ctx, "POST", path, conversion)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

func (c *Client) GetConversion(conversionId string) (Response, error) {
	path := converisionsUrl + "/" + conversionId

	ctx, cancel := context.WithTimeout(context.Background(), c.HttpClient.Timeout)
	defer cancel()

	response, err := c.sendRequest(ctx, "GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
