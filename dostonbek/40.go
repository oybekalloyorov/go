package main

import "fmt"

func main() {
	mp := map[string]int{"a": 2, "b": 3}
	mp2 := map[string]int{"b": 4, "c": 5}

	for s, w := range mp {
		text := ""
		son := 0
		for s2, w2 := range mp2 {
			if s == s2 {
				mp[s] = w + w2
			}
			text = s2
			son = w2
		}
		if son != 0 {
			mp[text] = son 
		}
	}
	fmt.Println(mp)
}