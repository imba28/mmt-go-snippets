package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan<- string) {
	for i := 0; ; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		c <- fmt.Sprintf("Worker %d: %d", id, i)
	}
}

// multiple channels are piped into one single channel
func fanIn(channels ...<-chan string) <-chan string {
	combinedChannel := make(chan string)

	for i := range channels {
		go func(c <-chan string) {
			for {
				combinedChannel <- <-c
			}
		}(channels[i])
	}

	return combinedChannel
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go worker(1, c1)
	go worker(2, c2)
	go worker(3, c3)

	for {
		select {
		case v := <-c1:
			fmt.Println("Worker 1 sends", v)
		case v := <-c2:
			fmt.Println("Worker 2 sends", v)
		case v := <-c3:
			fmt.Println("Worker 3 sends", v)
		}
	}

	/*for m := range fanIn(c1, c2, c3) {
		fmt.Println("A worker sends: ", m)
	}*/
}
