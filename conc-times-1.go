package main

import (
	"fmt"
	"time"
)

func main() {
	timestable(2)
}
func timestable(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)
		time.Sleep(100 * time.Millisecond)
	}
}
