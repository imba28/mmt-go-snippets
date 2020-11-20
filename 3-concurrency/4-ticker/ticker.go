package main

import (
	"fmt"
	"time"
)

// time.Tick - a value is sent to a channel periodically
func ticker() {
	ticker := time.Tick(time.Second)
	for v := range ticker {
		fmt.Println("TICKER: ", v)
	}
}

// time.After - after a certain time a value is sent to a channel
func timeout() {
	fmt.Println("TIMEOUT: start")
	timeout := time.After(2 * time.Second)
	// channel receives a value after 2 seconds
	<-timeout

	fmt.Println("TIMEOUT: end")
}

func main() {
	timeout()
	ticker()
}
