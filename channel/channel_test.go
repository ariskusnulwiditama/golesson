package channel

import (
	"fmt"
	"testing"
)

func TestChannel1(t *testing.T) {
	var message = make(chan string)

	for _, each := range []string{"one", "two", "three"} {
		go func(v string) {
			var data = fmt.Sprintf("the message: %s", v)
			message <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(message)
	}
}

func printMessage(input chan string) {
	fmt.Println(<-input)
}

func TestChannel2(t *testing.T) {
	var nums = []int{23, 54, 21, 11, 25}
	ch := make(chan float64)
	go GetAverage(nums, ch)
	ch2 := make(chan int)
	go GetMax(nums, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch:
			fmt.Printf("average= %.2f\n", avg)
		case max := <-ch2:
			fmt.Printf("max= %d\n", max)
		}
	}
}

func GetAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, v := range numbers {
		sum += v
	}
	ch <- float64(sum) / float64(len(numbers))
}

func GetMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, v := range numbers {
		if max < v {
			max = v
		}
	}

	ch <- max
}
