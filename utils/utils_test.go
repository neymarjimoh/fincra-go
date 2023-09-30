package utils

import (
	"testing"
)

type SubAccountBody struct {
	BusinessId string `json:"businessId"`
	Name       string `json:"name"`
}

func TestExcludeField(t *testing.T) {
	subAccount := SubAccountBody{
		BusinessId: "123",
		Name:       "John Doe",
	}

	newSubAccount := ExcludeField(subAccount, "BusinessId")

	if newSubAccount.BusinessId != "" {
		t.Errorf("Excluded field still present: BusinessId = %s", newSubAccount.BusinessId)
	}

	if newSubAccount.Name != subAccount.Name {
		t.Errorf("Name field is incorrect: expected %s, got %s", subAccount.Name, newSubAccount.Name)
	}

	if subAccount.BusinessId != "123" {
		t.Errorf("Original struct modified: BusinessId = %s", subAccount.BusinessId)
	}
}
