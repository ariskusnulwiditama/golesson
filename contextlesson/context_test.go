package contextlesson

import (
	"context"
	"runtime"
	"testing"

	_ "golesson/config"

	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	ctxBackground := context.Background()
	logrus.Info(ctxBackground)

	todo := context.TODO()
	logrus.Info(todo)
}

func TestWithValue(t *testing.T) {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b","B")
	ctxC := context.WithValue(ctxA, "c", "C")
	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")
	ctxF := context.WithValue(ctxC, "f", "F")
	logrus.Info("context A:", ctxA)
	logrus.Info("context B:",ctxB)
	logrus.Info("context C:",ctxC)
	logrus.Info("context D:",ctxD)
	logrus.Info("context E:",ctxE)
	logrus.Info("context F:",ctxF)

	logrus.Info(ctxF.Value("f"))
}

func CreateCounter() chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)
		counter := 1
		for {
			destination <-counter
			counter++
		}
	}()

	return destination
}

func TestCounter(t *testing.T) {
	logrus.Info("total goroutine: ", runtime.NumGoroutine())
}