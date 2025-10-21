package main

import "fmt"

func birlashtir(m, m2 map[string]int) map[string]int{
	result := make(map[string]int)

	for s, v := range m{
		result[s] = v
	}

	for s2, v2 := range m2{
		result[s2] += v2
	}

	return result
}

func main() {
	mp := map[string]int{"a": 2, "b": 3}
	mp2 := map[string]int{"b": 4, "c": 5, "b1": 4}

	res := birlashtir(mp, mp2)
	fmt.Println(res)
}