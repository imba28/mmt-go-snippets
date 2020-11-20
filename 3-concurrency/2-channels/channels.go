package main

import (
	"fmt"
	"time"
)

func greet(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("Hello there %d", i)
		time.Sleep(100 * time.Millisecond)
	}

	close(c)
}

func main() {
	c := make(chan string)

	go greet(c)

	for message := range c {
		fmt.Println(message)
	}

	fmt.Println("Exiting...")
}
