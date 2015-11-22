package main

import "fmt"

var (
	Version = "undefined"
	BuildTime = "undefined"
    GitHash = "undefined"
)

func main() {
    fmt.Printf("Version    : %s\n", Version)
    fmt.Printf("Git Hash   : %s\n", GitHash)
	fmt.Printf("Build Time : %s\n", BuildTime)
}
