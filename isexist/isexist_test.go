package isexist

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	myMap := map[string]int{
		"A": 230,
		"B": 904,
		"C": 339,
		"D": 223,
	}

	if _, ok := myMap["B"]; ok {
		fmt.Print("exist")
	}else{
		t.Fail()
		fmt.Print("not exist")
	}
}
