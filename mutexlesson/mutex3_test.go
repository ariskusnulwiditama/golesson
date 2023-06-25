package mutexlesson

import (
	"fmt"
	"sync"
	"testing"
)


var wg sync.WaitGroup
func TestWG(t *testing.T) {
	wg.Add(5)

	for i:=0; i<5; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("goroutine %d selesai\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("semua go routine selesai")
}
