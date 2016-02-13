package main

func bar(s string, i int) {
	panic(s)
}

func foo(s string) {
	bar(s, 1)
}

func main() {
	foo("ooh")
}
