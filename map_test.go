package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONDecode(t *testing.T) {
	jsonString := `{"id": "P001", "name": "Macbook Pro", "price" : 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])

}

func TestJSONEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Macbook Pro",
		"price": 20000000,
	}

	marshal, _ := json.Marshal(product)
	fmt.Println(string(marshal))
}
