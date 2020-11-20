package main

import (
	"fmt"
	"time"
)

func greet() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello there.")
	}
}

func main() {
	go greet()
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}
