package main

import (
	"fmt"
	"time"
)

func countTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	// blocking operation
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("we are done")
			break
		}

		fmt.Printf("received: %d\n", a)
	}
}

func main() {
	c := make(chan int)
	go countTen(c)
	receive(c)
}
