package main

import "fmt"

func printVals(a, b int) {
	

	ptrB := &a
	ptrA := &b
	fmt.Println("a:", a, "b:", b, "*a:", *ptrA, "*b:", *ptrB)
}

func main() {
	x := 3
	y := 7
	printVals(x, y)
}
