package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func Halo() {
	fmt.Println("halo")
}

func TestRoutine(t *testing.T) {
	go Halo()
	fmt.Println("dunia")
	time.Sleep(5 * time.Microsecond)
}
