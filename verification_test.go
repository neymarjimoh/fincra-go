package fincra

import (
	"fmt"
	"testing"
)

func TestVerifyBankAccount(t *testing.T) {
	t.Run("verify a bank account", func(t *testing.T) {
		client := defaultTestClient()

		request := VerifyBankAccountBody{
			AccountNumber: "0929292929",
			Type:          Nuban,
			BankCode:      "044",
			Iban:          "999",
		}

		resp, err := client.VerifyBankAccount(request)
		if err != nil {
			t.Errorf("error verifying bank account: %v", err)
		}

		fmt.Println(resp)

		// want := map[string]interface{}{
		// 	"success": true,
		// 	"message": "Account resolve successful",
		// }

		// got := make(map[string]interface{}, len(want))
		// for k, v := range resp {
		// 	if k == "message" || k == "success" {
		// 		got[k] = v
		// 	}
		// }

		// testEqual(t, got, want)
	})
}

func TestVerifyBVN(t *testing.T) {
	t.Run("verify a bvn", func(t *testing.T) {
		client := defaultTestClient()

		request := VerifyBVNBody{
			Bvn:      "09292929221",
			Business: "6457d39b12b4401f99a54772",
		}

		resp, err := client.VerifyBVN(request)
		if err != nil {
			t.Errorf("error verifying bvn: %v", err)
		}

		fmt.Println(resp)
	})
}
