package main

import "fmt"

// TODO: add fibonacci function

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
