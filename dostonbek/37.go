package main

import "fmt"

func main(){
	text := "hello"

	counts := make(map[rune]int)

	for i := 0; i < len(text); i++ {
        fmt.Printf("Index: %d, Char: %c\n", i, text[i])
    }

	for _, v := range text{
		counts[v]++
	}

	for key, value := range counts {
		fmt.Printf("%c: %d\n", key, value)	
	}



}