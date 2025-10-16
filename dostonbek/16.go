package main

import "fmt"

func Ekub(a, b int) {
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println(a)
}

func main() {
	Ekub(24, 36)
}