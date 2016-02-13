package main

import "sync"

func hang(i *int) {
	var deadlock chan bool
	<-deadlock
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		j := new(int)
		*j = i
		go func(k *int) {
			defer wg.Done()
			hang(k)
		}(j)
	}
	wg.Wait()
}
