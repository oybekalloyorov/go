package main

import (
	"fmt"
)
func printItem[T any](item, defaultValue T)(T, T) {
	return item, defaultValue
}

func main() {
	num1, num2 := printItem(10, 20)
	str1, str2 := printItem[string]("Hello", "World")
	bool1, bool2 := printItem[bool](true, false)

	fmt.Println(num1, num2)
	fmt.Println(str1, str2)	
	fmt.Println(bool1, bool2)
	
}