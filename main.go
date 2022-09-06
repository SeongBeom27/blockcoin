package main

import (
	"fmt"
	"time"
)

func countTen(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf(">> sending %d <<\n", i)
		c <- i
		fmt.Printf(">> sent %d <<\n", i)
	}
	close(c)
}

func receive(c <-chan int) {
	// blocking operation
	for {
		time.Sleep(10 * time.Second)
		a, ok := <-c
		if !ok {
			fmt.Println("we are done")
			break
		}
		fmt.Printf("|| received: %d ||\n", a)
	}
}

func main() {
	c := make(chan int, 5)
	go countTen(c)
	receive(c)
}
