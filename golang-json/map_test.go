package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"1","name":"Product 1","ImageURL":123}`
	jsonBytes := []byte(jsonString)
	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["ImageURL"])
}
