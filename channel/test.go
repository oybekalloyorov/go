package main

import (
	"fmt"
	"time"
	"math/rand"
)

func search(term string, urls []string, results chan<- string) {
	for _, url := range urls {
		// Simulate a search operation with a random delay
		time.Sleep(time.Millisecond * 1000)
		if rand.Float32() < 0.3 { // 30% chance to find the term
			results <- fmt.Sprintf("topildi '%s' in %s", term, url)
			return
		}
	}
}
func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
		"https://google.com",
}
	results := make(chan string)

	go search("Go", urls, results)
	go search("Python", urls, results)
	go search("Java", urls, results)

	
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
	// for result := range results {
	// 	fmt.Println(result)
	// }
	close(results)
}