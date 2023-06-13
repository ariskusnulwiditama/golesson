package containerlesson

import (
	"container/list"
	"fmt"
	"log"
	"testing"
)

func TestContainerLesson(t *testing.T) {
	myList := list.New()
	myList.PushBack("one")
	myList.PushBack("two")
	myList.PushBack("three")
	myList.PushBack("four")
	myList.PushBack("five")

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	} 
	search, elem := SearchList(9, myList)
	log.Println("element in index ", search, ": ", elem)
}

func SearchList(search int, listP *list.List) (int, interface{}){
	currenIndex := 1
	var i interface{}
	for e := listP.Front(); e != nil; e = e.Next() {
		if currenIndex == search {
			i = e.Value
			break
		}else{
			return search, "not exist"
		}
		
	}
	return search, i
}