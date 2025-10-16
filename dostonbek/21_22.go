package main

import "fmt"

func ptrChange(index, value int, a *[4]int)( *[4]int) {
	
	(*a)[index] = value
	return a
}

func main() {
	arr := [4]int{1, 2, 3, 4}
	index := 2
	value := 9
	fmt.Println(ptrChange(index, value, &arr))
}