package utils

import (
	"reflect"
)

func ExcludeField[T interface{}](input T, fieldName string) T {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() != reflect.Struct {
		panic("Input is not a struct")
	}

	outputValue := reflect.New(inputValue.Type()).Elem()

	for i := 0; i < inputValue.NumField(); i++ {
		field := inputValue.Type().Field(i)
		fieldValue := inputValue.Field(i)

		if field.Name != fieldName {
			outputValue.FieldByName(field.Name).Set(fieldValue)
		}
	}

	return outputValue.Interface().(T)
}
