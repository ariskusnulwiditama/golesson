package parse

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	myMap := map[string] interface{}{
		"name": "dean",
		"age": 44,
		"isLive": true,
	}

	var mySlice []interface{}
	for _, v := range myMap {
		mySlice = append(mySlice, v)
	}
	log.Println(mySlice)
}

func TestMapToSlice(t *testing.T) {
	mySlice := []interface{} {
		map[string] interface{}{
			"name": "dean",
			"age": 34,
			"isLive": true,
		},
	}

	res := ParseToMap(mySlice)
	log.Println(res)
}

func ParseToMap(input []interface{}) map[string] interface{} {
	result := make(map[string] interface{})
	for _, v := range input {
		if kv, ok := v.(map[string] interface{}); ok {
			for k, v := range kv {
				result[k] = v
			}
		}
	}
	return result
}