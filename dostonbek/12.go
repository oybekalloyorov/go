package main

import (
	"fmt"
	"math"
)

func daraja_rec(a, b int) int {
	if b == 0 {
		return 1
	}
	if b < 0 {
		return 0
	}
	return int(math.Pow(float64(a), float64(b)))
}

func daraja(a, b int) {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	fmt.Println(result)
}

func main() {
	daraja(2, 3)
	fmt.Println(daraja_rec(2, 3))
}