package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	// CasualAddress used non-secured address to listen
	CasualAddress = ":80"
	// SecureAddress used for secured address to listen
	SecureAddress = ":443"
)

func init() {
	go signalHandler(make(chan os.Signal, 1))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	})

	var servers sync.WaitGroup

	launch("HTTP", CasualAddress, &servers, func(address string) error {
		return http.ListenAndServe(address, nil)
	})

	launch("HTTPS", SecureAddress, &servers, func(address string) error {
		return http.ListenAndServeTLS(address, "cert.pem", "key.pem", nil)
	})

	fmt.Println("Ready for connections");

	servers.Wait()
}

func launch(name string, address string, wg *sync.WaitGroup, f func(address string) error) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println(name, "binding to", address);

		if e := f(address); e != nil {
			fmt.Println(name, "->", e)

			syscall.Kill(syscall.Getpid(), syscall.SIGABRT)
		}
	}()
}

func signalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)

	for s := <-c; ; s = <-c {
		switch s {
		case syscall.SIGABRT:
			fmt.Println("abnormal exit")
			os.Exit(1)
		case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("clean shutdown")
			os.Exit(0)
		}
	}
}
