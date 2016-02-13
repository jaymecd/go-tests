package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type mutexCounter struct {
	sync.RWMutex
	x int64
}

func (c *mutexCounter) Add(x int64) {
	c.Lock()
	c.x += x
	c.Unlock()
}

func (c *mutexCounter) Value() (x int64) {
	c.RLock()
	x = c.x
	c.RUnlock()
	return
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()

	fmt.Printf("Proc: %d\nCPUs: %d\n", maxProcs, numCPU)

	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	fmt.Printf("Parallelism: %d\n", MaxParallelism())

	counter := mutexCounter{}

	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value())

}
