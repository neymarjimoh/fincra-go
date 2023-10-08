package fincra

import (
	"fmt"
	"testing"
)

func TestCreateVirtualAccount(t *testing.T) {
	request := CreateVirtualAccountDto{
		Currency:    "NGN",
		UtilityBill: "https://www.planetware.com/wpimages/2020/02/france-in-pictures-beautiful-places-to-photograph-eiffel-tower.jpg",
		AccountType: "individual",
		KYCInformation: KYCInformationDto{
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "abc@abc.com",
			BusinessName: "JohnDoe",
		},
		Channel: "providus",
	}

	t.Run("create virtual account", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.CreateVirtualAccount(request)
		if err != nil {
			t.Errorf("error creating virtual account: %v", err)
		}

		fmt.Println(resp)
	})
}

func TestListVirtualAccounts(t *testing.T) {
	t.Run("list virtual accounts", func(t *testing.T) {
		client := defaultTestClient()

		options := Options{
			Currency: "EUR",
		}

		resp, err := client.ListVirtualAccounts(options)
		if err != nil {
			t.Errorf("error listing virtual accounts: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Merchant virtual accounts fetched successfully",
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

func TestListVirtualAccountRequests(t *testing.T) {
	t.Run("list virtual account requests", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.ListVirtualAccountRequests()
		if err != nil {
			t.Errorf("error listing virtual account requests: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Virtual account requests fetched successfully",
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

func TestListVirtualAccountByCurrency(t *testing.T) {
	t.Run("Fetch a virtual account by currency", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.ListVirtualAccountByCurrency("EUR")
		if err != nil {
			t.Errorf("error fetching a virtual account by currency: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Merchant virtual accounts fetched successfully",
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

func TestListVirtualAccountByBvn(t *testing.T) {
	t.Run("Fetch a virtual account by BVN", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.ListVirtualAccountByBvn("0123456789", "6457d39b12b4401f99a54772")
		if err != nil {
			t.Errorf("error fetching a virtual account by BVN: %v", err)
		}

		fmt.Println(resp)
	})
}

func TestListVirtualAccount(t *testing.T) {
	t.Run("Fetch a virtual account", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.ListVirtualAccount("6457d39b12b4401f99a54772")
		if err != nil {
			t.Errorf("error fetching a virtual account: %v", err)
		}

		fmt.Println(resp)
	})
}
