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

func IsEmpty(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	if value.Kind() != reflect.Struct {
		panic("Input is not a struct")
	}

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		zeroValue := reflect.Zero(fieldValue.Type())

		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
			return false
		}
	}

	return true
}
