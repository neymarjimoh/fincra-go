package fincra

import (
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

	response, err := c.sendRequest("GET", path, nil)

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

	response, err := c.sendRequest("POST", path, conversion)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}

func (c *Client) GetConversion(conversionId string) (Response, error) {
	path := converisionsUrl + "/" + conversionId

	response, err := c.sendRequest("GET", path, nil)

	_ = json.Unmarshal(response, &jsonResponse)

	return jsonResponse, err
}
