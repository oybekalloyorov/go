package main

import (
	"fmt"
	"time"
	// "time"
)

// func generator(ch chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i // Send the number to the channel
// 	}
// 	close(ch) // Close the channel when done
// }
func main() {
	// vaziva 1

	// ch := make(chan string, 1) // Create a channel to pass ints
	// ch <- "Salom Dunyo" // Send a value to the channel
	// msg := <-ch // Receive a value from the channel
	// fmt.Println(msg) // Print the received value

	// vaziva 2
	// now := time.Now()
	// ch := make(chan int)
	// go generator(ch)

	// sum := 0
	// for num := range ch {
	// 	sum += num
	// }

	// fmt.Println("Yig'indi:", sum, "Vaqt:", time.Since(now))

	// vaziva 3 Azim polat 7:40:00
	// c := make(chan int, 5) // Create a channel to pass ints
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		c <- i
	// 	}
	// }
	// close(c) // Close the channel when done
	// for num := range c {
	// 	fmt.Println(num)
	// }

	// vaziva 4
	tick := time.Tick(500 * time.Millisecond)
	boom := time.Tick(2 * time.Second)
	timeout := time.After(5 * time.Second)
	
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom")
		case <-timeout:
			fmt.Println("Dastur tugadi")
			return
		}	
	}




}