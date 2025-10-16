package main

import "fmt"

func Ekuk(a, b int) {
	originalA, originalB := a, b
	for b != 0 {
		a, b = b, a%b
		fmt.Println("a->",a, " b->",b)
	}
	fmt.Println((originalA * originalB) / a)
}

func main() {
	Ekuk(6, 4)
}