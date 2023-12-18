package fincra

import (
	"context"
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

// ListChargeBacks lists all the chargebacks incurred on the business
func (c *Client) ListChargeBacks(ctx context.Context, businessId string) (Response, error) {
	path := chargebacksUrl + "?business=" + businessId

	return c.sendRequest(ctx, "GET", path, nil)
}

// AcceptChargeBack accepts a chargeback
func (c *Client) AcceptChargeBack(ctx context.Context, body *AcceptChargeBackDto) (Response, error) {
	if body.BusinessId == "" {
		return Response{}, errors.New("businessId is required to accept the chargeback")
	}

	if body.ChargeBackId == "" {
		return Response{}, errors.New("chargeBackId is required to update the chargeback")
	}

	path := chargebacksUrl + "/" + body.ChargeBackId + "/accept?business=" + body.BusinessId

	return c.sendRequest(ctx, "PATCH", path, nil)
}

// RejectChargeBack rejects a chargeback
func (c *Client) RejectChargeBack(ctx context.Context, body *RejectChargeBackDto) (Response, error) {
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

	return c.sendRequest(ctx, "PATCH", path, body)
}
