package main

import "fmt"

func pointerPlus(x *int) int{
	*x ++
	return *x
}

func main() {
	x := 10;
	pointerPlus(&x)
	fmt.Println(x)
}