package main

import (
	"fmt"
	"net/http"
)

const (
	CASUAL_ADDRESS = ":80"
	SECURE_ADDRESS = ":443"
)

func main() {
	message := "hello world"

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})

	go func() {
		if e := http.ListenAndServe(CASUAL_ADDRESS, nil); e != nil {
			panic(e)
		}
	}()

    go func() {
        if e := http.ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil); e != nil {
			panic(e)
		}
    }()

	for {}
}
