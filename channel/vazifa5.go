package main

import (
	"fmt"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- task * task
	}
}

func main() {
	tasks := make(chan int)
	results := make(chan int)

	// Ishchilar soni
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		id := i
		go worker(id, tasks, results)
	}
	// Vazifalarni yuborish
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)
	// Natijalarni qabul qilish
	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}
}
