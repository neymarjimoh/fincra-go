package utils

import (
	"testing"
)

type SubAccountBody struct {
	BusinessId string `json:"businessId"`
	Name       string `json:"name"`
}

type TestStruct1 struct {
	StrValue string
	IntValue int
}
type TestStruct2 struct {
	BoolValue bool
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

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		expect bool
	}{
		{"Empty struct", TestStruct1{}, true},
		{"Non-empty struct", TestStruct1{"hello", 42}, false},
		{"Empty struct with bool", TestStruct2{}, true},
		{"Non-empty struct with bool", TestStruct2{true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isEmptyResult := IsEmpty(tt.input)
			if isEmptyResult != tt.expect {
				t.Errorf("Expected isEmpty(%v) to be %v, but got %v", tt.input, tt.expect, isEmptyResult)
			}
		})
	}
}
