package mutexlesson

import (
	"sync"
	"testing"
)

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) Value() int {
	return c.val
}

func TestMutexRace(t *testing.T) {
	var wg sync.WaitGroup
	var meter  counter
	
	for i:=0; i<1000; i++ {
		wg.Add(1)

		go func()  {
			for k:=0; k<1000; k++ {
				meter.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
