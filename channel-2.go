// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.

package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan int) {

	for i := 0; i < 5; i++ {
		fmt.Printf("%d) working...", i)
		time.Sleep(time.Second)
		fmt.Println("done")

		// Send a value to notify that we're done.
		done <- i
	}
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan int, 1)

	fmt.Println("> start")

	go worker(done)

	for i := 0; i < 5; i++ {
		// Block until we receive a notification from the
		// worker on the channel.
		x := <-done
		fmt.Printf("\n>> %d) got\n", x)
	}

	fmt.Println("> stop")
}
