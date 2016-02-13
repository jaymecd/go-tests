package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func こんにちは_Gophers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "こんにちは  Gophers!\nnow: %s\n", time.Now())
}

func main() {
	http.HandleFunc("/", こんにちは_Gophers)
	err := http.ListenAndServe("localhost:12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
