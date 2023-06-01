package fincra

import (
	"fmt"
	"testing"
)

func TestCreateBeneficiary(t *testing.T) {
	createBeneficiary := &CreateBeneficiaryBody{
		BusinessId:        "6457d39b12b4401f99a54772",
		FirstName:         "Test1",
		LastName:          "Test2",
		Email:             "abc@abc.com",
		PhoneNumber:       "09090909091",
		AccountHolderName: "Test1 test2 Test3",
		Bank: Bank{
			Name:     "Wema Bank",
			Code:     "044",
			SortCode: "928927",
			Address: Address{
				Country: "GB",
				City:    "London",
				Street:  "Test Street",
				Zip:     "123",
				State:   "Osapa London",
			},
		},
		Type:               Individual,
		Currency:           "GBP",
		PaymentDestination: CryptoWallet,
		UniqueIdentifier:   "1",
		DestinationAddress: "Osapa London",
	}

	t.Run("create beneficiary", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.CreateBeneficiary(createBeneficiary)
		if err != nil {
			t.Errorf("error creating beneficiary: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "You have successfully created a beneficiary",
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

func TestGetBeneficiaries(t *testing.T) {
	getBeneficiary := &GetAllBeneficiariesParams{
		BusinessId: "6457d39b12b4401f99a54772",
		Page:       "1",
		PerPage:    "10",
	}

	t.Run("get all beneficiaries", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.GetAllBeneficiaries(getBeneficiary)
		if err != nil {
			t.Errorf("error getting all beneficiaries: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Beneficiaries fetched successfully",
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

func TestGetBeneficiary(t *testing.T) {
	getABeneficiary := &GetBeneficiaryParams{
		BusinessId:    "6457d39b12b4401f99a54772",
		BeneficiaryId: "646db15a8cdec23981165184",
	}

	t.Run("get beneficiary", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.GetBeneficiary(getABeneficiary)
		if err != nil {
			t.Errorf("error getting the specified beneficiary: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Beneficiary fetched successfully",
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

func TestUpdateBeneficiary(t *testing.T) {
	updateBeneficiary := &UpdateBeneficiaryBody{
		BusinessId:    "6457d39b12b4401f99a54772",
		BeneficiaryId: "646db15a8cdec23981165184",
		FirstName:     "Tester 3",
		Type:          Individual,
	}

	t.Run("update beneficiary", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.UpdateBeneficiary(updateBeneficiary)
		if err != nil {
			t.Errorf("error updating beneficiary: %v", err)
		}

		fmt.Println(resp)
	})
}

func TestDeleteBeneficiary(t *testing.T) {
	getABeneficiary := &GetBeneficiaryParams{
		BusinessId:    "6457d39b12b4401f99a54772",
		BeneficiaryId: "646db15a8cdec23981165184",
	}

	t.Run("delete beneficiary", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.DeleteBeneficiary(getABeneficiary)
		if err != nil {
			t.Errorf("error deleting the specified beneficiary: %v", err)
		}

		fmt.Println(resp)
	})
}
