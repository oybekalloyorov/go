package main

import (
	"fmt"
)

func PrintSlice[T any](s []T) []T {
	for _, v := range s {
		fmt.Println(v)
	}
	return s
}
func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	PrintSlice(intSlice)

	strSlice := []string{"Hello", "World", "Go", "Generics"}
	PrintSlice(strSlice)

}