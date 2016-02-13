package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Counter struct {
	sync.RWMutex
	value int
}

func (c *Counter) increment() {
	c.Lock()
	defer c.Unlock()

	c.value++
}

func (c *Counter) reset() {
	c.Lock()
	defer c.Unlock()

	c.value = 0
}

func (c *Counter) get() int {
	c.RLock()
	defer c.RUnlock()

	return c.value
}

var c = new(Counter)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/reset", reset)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	c.increment()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handler echoes the Path component of the requested URL.
func reset(w http.ResponseWriter, r *http.Request) {
	c.reset()

	fmt.Fprintf(w, "reseted")
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	count := c.get()

	fmt.Fprintf(w, "Count %d\n", count)
}
