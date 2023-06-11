package fincra

import (
	"fmt"
	"testing"
)

func TestGetBusinessConversions(t *testing.T) {
	t.Run("get business conversions", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.GetBusinessConversions("6457d39b12b4401f99a54772")
		if err != nil {
			t.Errorf("error getting conversions: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Conversions fetched successfully",
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

func TestCreateConversion(t *testing.T) {
	createConversion := &CreateConversionBody{
		BusinessId:     "6457d39b12b4401f99a54772",
		QuoteReference: "randomTextFromQuote",
	}

	t.Run("create conversion", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.CreateConversion(createConversion)
		if err != nil {
			t.Errorf("error creating conversion: %v", err)
		}

		fmt.Println(resp)
	})
}

func TestGetConversion(t *testing.T) {
	t.Run("get a conversion", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.GetConversion("12345678")
		if err != nil {
			t.Errorf("error getting a conversion: %v", err)
		}

		fmt.Println(resp)
	})
}
