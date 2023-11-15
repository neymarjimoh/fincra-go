package fincra

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateQuote(t *testing.T) {
	createQuote := &CreateQuoteBody{
		Action:              "send",
		TransactionType:     "conversion",
		FeeBearer:           "business",
		PaymentDestination:  FliqPayWallet,
		BeneficiaryType:     Individual,
		Business:            "6457d39b12b4401f99a54772",
		Amount:              "150",
		DestinationCurrency: "USD",
		SourceCurrency:      "NGN",
	}

	t.Run("create quote", func(t *testing.T) {
		client := defaultTestClient()
		ctx := context.Background()

		resp, err := client.CreateQuote(ctx, createQuote)
		if err != nil {
			t.Errorf("error creating quote: %v", err)
		}

		fmt.Println(resp)
	})
}
