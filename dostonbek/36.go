package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	text := "go is fun and go is fast"

	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, w := range words {
		counts[w]++
	}

	// Tartib uchun keys yig‘amiz
    keys := make([]string, 0, len(counts))
    for k := range counts {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    for _, k := range keys {
        fmt.Printf("%s -> %d ", k, counts[k])
    }	
}