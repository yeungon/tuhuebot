package helpers

import (
	"fmt"
	"reflect"
)

func PrintStruct(v interface{}) {
	// Get the value and type of the struct
	value := reflect.ValueOf(v)
	typeOfValue := value.Type()

	// Check if the input is a pointer to a struct
	if value.Kind() == reflect.Ptr {
		value = value.Elem() // Get the value pointed to by the pointer
		typeOfValue = value.Type()
	}

	// Ensure the value is a struct
	if value.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct.")
		return
	}

	// Iterate over the struct fields
	for i := 0; i < value.NumField(); i++ {
		field := typeOfValue.Field(i)            // Get the struct field's metadata
		fieldValue := value.Field(i).Interface() // Get the value of the field

		// Print the field name, type, and value
		fmt.Printf("%s (%s): %v\n", field.Name, field.Type, fieldValue)
	}
}
