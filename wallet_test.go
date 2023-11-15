package fincra

import (
	"context"
	"fmt"
	"testing"
)

func TestGetWallets(t *testing.T) {
	t.Run("fetch wallets", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		resp, err := client.ListWallets(ctx, "6457d39b12b4401f99a54772")
		if err != nil {
			t.Errorf("error fetching wallets info: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Wallets fetched successfully",
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

func TestGetWallet(t *testing.T) {
	t.Run("fetch wallet", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		resp, err := client.ListWallet(ctx, "66433")
		if err != nil {
			t.Errorf("error fetching wallet info: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Wallet fetched successfully",
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

func TestGetWalletLogs(t *testing.T) {
	t.Run("fetch wallet logs", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		request := LogsDto{
			Business: "6457d39b12b4401f99a54772",
			Action:   Credit,
			Page:     "2",
			PerPage:  "10",
			Amount:   "500",
		}

		resp, err := client.ListWalletLogs(ctx, request)
		if err != nil {
			t.Errorf("error fetching wallet info: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Wallet logs fetched successfully",
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
