package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("tes")
	logJson(1)
	logJson(true)
	logJson([]string{"a", "b", "c"})
}	