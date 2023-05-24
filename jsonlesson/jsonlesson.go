package jsonlesson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}
type Peoople struct {
	Name    string    `json:"name"`
	Address []Address `json:"address"`
	Married bool      `json:"married"`
	Balance int32     `json:"balance"`
	Hobbies []string  `json:"hobbies"`
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func encodeJson() ([]byte, error) {
	people := Peoople{
		Name: "John",
		Address: []Address{
			{Street: "st witch",
				Country:    "US",
				PostalCode: "7654",
			},
			{Street: "st wesmingham",
				Country:    "US",
				PostalCode: "54376",
			},
		},
		Married: true,
		Balance: 30000,
		Hobbies: []string{"game", "reading", "climbing"},
	}

	res, err := json.Marshal(people)
	return res, err
}

func decodeJson() {
	jsonText := `{"name":"John","address":[{"Street":"st witch","Country":"US","PostalCode":"7654"},{"Street":"st wesmingham","Country":"US","PostalCode":"54376"}],"married":true,"balance":30000,"hobbies":["game","reading","climbing"]}`
	jsonBytext := []byte(jsonText)

	people := Peoople{}
	err := json.Unmarshal(jsonBytext, &people)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name: ", people.Name)
	fmt.Println("Address: ", people.Address)
	fmt.Println("Balance: ", people.Balance)
	fmt.Println("Married: ", people.Married)
	fmt.Println("Hobbies: ", people.Hobbies)
}

func mapJsonDecode() {
	jsonString := `{"id":"ap001", "name":"apple pro", "price": 4000000}`
	jsonBytes := []byte(jsonString)
	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}

func mapJsonEncode() {
	product := map[string]interface{}{
		"id":    "apl223",
		"name":  "apple luminar",
		"price": 12000000,
	}
	res, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}

func streamDecoder() {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)
	cust := &Customer{}
	decoder.Decode(cust)

	fmt.Println(cust)
}

func streamEncoder() {
	writer, err := os.Create("customer_out.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	
	cust := Customer {
		FirstName: "Michael",
		MiddleName: "Valos",
		LastName: "Quincy",
	}

	encoder.Encode(cust)

}