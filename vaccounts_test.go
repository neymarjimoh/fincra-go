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
		KYCInformation: KYCInformation{
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "abc@abc.com",
			BusinessName: "JohnDoe",
		},
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
