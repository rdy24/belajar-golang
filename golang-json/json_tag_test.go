package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL int    `json:"img_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "Product 1",
		ImageURL: 123,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"1","name":"Product 1","ImageURL":123}`
	jsonBytes := []byte(jsonString)
	var product Product
	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
