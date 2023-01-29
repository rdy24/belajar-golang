package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreaming(t *testing.T) {
	reader, _ := os.Open("test.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}

func TestStreamingEncode(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
	}

	encoder.Encode(customer)

}
