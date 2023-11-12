package fincra

import (
	"errors"
)

const chargebacksUrl = "/collections/chargebacks"

type AcceptChargeBackDto struct {
	ChargeBackId string `json:"chargeBackId"`
	BusinessId   string `json:"businessId"`
}

type RejectChargeBackDto struct {
	ChargeBackId string `json:"chargeBackId"`
	BusinessId   string `json:"businessId"`
	Reason       string `json:"business_reject_reason"`
}

func (c *Client) ListChargeBacks(businessId string) (Response, error) {
	path := chargebacksUrl + "?business=" + businessId

	return c.sendRequest("GET", path, nil)
}

func (c *Client) AcceptChargeBack(body *AcceptChargeBackDto) (Response, error) {
	if body.BusinessId == "" {
		return Response{}, errors.New("businessId is required to accept the chargeback")
	}

	if body.ChargeBackId == "" {
		return Response{}, errors.New("chargeBackId is required to update the chargeback")
	}

	path := chargebacksUrl + "/" + body.ChargeBackId + "/accept?business=" + body.BusinessId

	return c.sendRequest("PATCH", path, nil)
}

func (c *Client) RejectChargeBack(body *RejectChargeBackDto) (Response, error) {
	if body.BusinessId == "" {
		return Response{}, errors.New("businessId is required to reject the chargeback")
	}

	if body.ChargeBackId == "" {
		return Response{}, errors.New("chargeBackId is required to reject the chargeback")
	}

	if body.Reason == "" {
		return Response{}, errors.New("reason is required to reject the chargeback")
	}

	path := chargebacksUrl + "/" + body.ChargeBackId + "/reject?business=" + body.BusinessId

	return c.sendRequest("PATCH", path, body)
}
