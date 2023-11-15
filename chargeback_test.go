package fincra

import (
	"context"
	"fmt"
	"testing"
)

const businessId = "6457d39b12b4401f99a54772"

func TestListChargebacks(t *testing.T) {
	t.Run("list business chargebacks", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		resp, err := client.ListChargeBacks(ctx, businessId)
		if err != nil {
			t.Errorf("error lsiting all chargebacks: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Chargebacks fetched successfully",
		}

		got := make(map[string]interface{}, len(want))
		for k, v := range resp {
			if k == "message" || k == "success" {
				got[k] = v
			}
		}

		testEqual(t, got, want)
	})
}

func TestAcceptChargeBack(t *testing.T) {
	acceptChargeBack := &AcceptChargeBackDto{
		BusinessId:   businessId,
		ChargeBackId: "7171892", // random id
	}

	t.Run("accept chargeback", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		resp, err := client.AcceptChargeBack(ctx, acceptChargeBack)
		if err != nil {
			t.Errorf("error accepting chargeback: %v", err)
		}

		fmt.Println(resp)
	})
}

func TestRejectChargeBack(t *testing.T) {
	rejectChargeBack := &RejectChargeBackDto{
		BusinessId:   businessId,
		ChargeBackId: "7171892", // random id
		Reason:       "no reason",
	}

	t.Run("reject chargeback", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		resp, err := client.RejectChargeBack(ctx, rejectChargeBack)
		if err != nil {
			t.Errorf("error rejecting chargeback: %v", err)
		}

		fmt.Println(resp)
	})
}
