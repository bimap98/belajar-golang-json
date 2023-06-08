package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Eko",
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Eko",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "8888",
			},
			{
				Street:     "Jalan Masih Dibangun",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"8888"},{"Street":"Jalan Masih Dibangun","Country":"Indonesia","PostalCode":"9999"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	address := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "8888",
		},
		{
			Street:     "Jalan Masih Dibangun",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
	}

	bytes, _ := json.Marshal(address)
	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"8888"},{"Street":"Jalan Masih Dibangun","Country":"Indonesia","PostalCode":"9999"}]`
	jsonBytes := []byte(jsonString)

	address := &[]Address{}

	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)
}
