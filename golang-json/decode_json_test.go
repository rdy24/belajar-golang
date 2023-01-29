package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Doe"}`
	jsonBytes := []byte(jsonString)
	var customer Customer
	json.Unmarshal(jsonBytes, &customer)
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}