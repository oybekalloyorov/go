package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; i < 1_000_000; i++ {
		ch <- i
	}
}

func main() {
	now := time.Now()
	ch := make(chan int)

	go sendData(ch)
	value := 0
	for i := 0; i < 1_000_000; i++ {
		value = <-ch
		// fmt.Println("Received:", val)
	}
	fmt.Println("Finished receiving 1,000,000 values", value)
	fmt.Println("Time taken:", time.Since(now))
}