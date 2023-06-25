package goroutine

import (
	"fmt"
	"log"
	"math/rand"
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

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
}

func TestSayGreeting(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello", 10)
	time.Sleep(2 * time.Second)
}
