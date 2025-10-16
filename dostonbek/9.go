package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := n * factorial(n-1)
	// fmt.Println(result)
	return result
}

func main() {
	fmt.Println(factorial(4))
}