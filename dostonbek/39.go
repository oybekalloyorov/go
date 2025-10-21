package main

import "fmt"

func main() {
	mp := map[string]string{"Ali": "A", "Bob": "B", "Eve": "A"}

	mp2 := make(map[string][]string)

	for s, v := range mp {
		mp2[v] = append(mp2[v], s)
	}

	
	fmt.Println(mp2)
	

}