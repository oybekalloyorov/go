package main

import "fmt"


func tub_son(n int) {
	for i := 2; i <= n; i++ {
		tub := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				tub = false
				break
			}
		}
		if tub {
			fmt.Println(i)
		}
	}
}

func tubmi(n int) string {
	if n < 2 {
		return "tub emas"
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return "tub emas"
		}
	}
	return "tub son"
}

func main() {
	tub_son(20)

	fmt.Println("19 ", tubmi(19))
	fmt.Println("20 ", tubmi(20))
}