package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Hobbies  []string
	Addresses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
