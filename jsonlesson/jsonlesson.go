package jsonlesson

import (
	"encoding/json"
	"fmt"
)

type Peoople struct {
	Name    string   `json:"name"`
	Address []Address   `json:"address"`
	Married bool     `json:"married"`
	Balance int32    `json:"balance"`
	Hobbies []string `json:"hobbies"`
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
		Name:    "John",
		Address: []Address{
			{Street: "st witch",
			Country: "US",
			PostalCode: "7654",
			},
			{Street: "st wesmingham",
			Country: "US",
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

	fmt.Println("Name: ",people.Name)
	fmt.Println("Address: ", people.Address)
	fmt.Println("Balance: ", people.Balance)
	fmt.Println("Married: ", people.Married)
	fmt.Println("Hobbies: ", people.Hobbies)
}
