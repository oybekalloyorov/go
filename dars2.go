package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func main(){
	var arr [3]int = [3]int{1, 2, 3}

	fmt.Println("array:", arr, "len:", len(arr))

	s := []int{1,2,3,4}
	fmt.Println("slice", s, "len:", len(s), "cap:", cap(s))
	
	s = append(s, 5)
	fmt.Println("after append", s)
	// fmt.Println("slice", s, "len:", len(s), "cap:", cap(s))

	s2 := s[1:4]
	fmt.Println("s[1:4] =", s2)
	
	 // make va copy
    s3 := make([]int, 0, 10) // uzunligi 0, sig'imi 10
    s3 = append(s3, 100, 200)
    s4 := make([]int, len(s3))
    copy(s4, s3)
    fmt.Println(s3, s4)

	//Map
	m := map[string]int{"one":1, "two":2}
	m["three"] = 3

	if v, ok := m["two"]; ok {
		fmt.Println("two:", v)
	}

	delete(m, "one")
	fmt.Println("map:", m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Strings va runes (UTF-8)
	satr := "Salom 世界"
	fmt.Println("bytes:", len(satr))
	fmt.Println("runes:", utf8.RuneCountInString(satr))
	fmt.Println(strings.ToUpper(satr))
	fmt.Println(strings.Fields("one two three")) // ["one","two","three"]

}