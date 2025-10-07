package main

import (
	"fmt"
	"time"
)

func main() {
	tick := make(chan string)
	boom := make(chan string)

	// tick gorutina
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			tick <- "tick"
		}
	}()

	// boom gorutina
	go func() {
		for {
			time.Sleep(2 * time.Second)
			boom <- "BOOM!"
		}
	}()

	timeout := time.After(5 * time.Second)

	for {
		select {
		case msg := <-tick:
			fmt.Println(msg)
		case msg := <-boom:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("⏰ Dastur tugadi.")
			return
		}
	}
}
