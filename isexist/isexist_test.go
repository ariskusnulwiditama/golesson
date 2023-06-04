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

func TestSliceExist(t *testing.T) {
	mySlice := []interface{}{
		123, "one", true, 123.55,
	}

	for _, v := range mySlice {
		if v ==1239 {
			fmt.Println("there")
			return
		}else{
			fmt.Println("not there")
			return
		}
	}
}