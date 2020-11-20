package main

import (
	"fmt"
	"time"
)

func person(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	start := time.Now()

	const n = 100000
	lastPerson := make(chan int)
	personToTheLeft := lastPerson

	for i := 0; i < n; i++ {
		personToTheRight := make(chan int)
		go person(personToTheLeft, personToTheRight)
		personToTheLeft = personToTheRight
	}

	personToTheLeft <- 1

	fmt.Println(<-lastPerson)
	fmt.Println(time.Since(start))
}
