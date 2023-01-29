package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
		Hobbies: []string{"a", "b", "c"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Doe","Hobbies":["a","b","c"]}`
	jsonBytes := []byte(jsonString)
	var customer Customer
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
		Hobbies: []string{"a", "b", "c"},
		Addresses: []Address{
			{
				Street: "Jl. ABC",
				Country: "Indonesia",
				PostalCode: "12345",
			},
			{
				Street: "Jl. DEF",
				Country: "Indonesia",
				PostalCode: "12345",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecodeComplex(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Doe","Hobbies":["a","b","c"],"Addresses":[{"Street":"Jl. ABC","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. DEF","Country":"Indonesia","PostalCode":"12345"}]}`
	jsonBytes := []byte(jsonString)
	var customer Customer
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestJSONArrayDecodeArray(t *testing.T) {
	jsonString := `[{"Street":"Jl. ABC","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. DEF","Country":"Indonesia","PostalCode":"12345"}]`
	jsonBytes := []byte(jsonString)
	var addresses []Address
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addresses)
}