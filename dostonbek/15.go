package main

import "fmt"

func palindromSon(n int) {
	original := n
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	if original == reversed {
		fmt.Println("Palindrom son -> ", original)
	}else {
		fmt.Println("Palindrom son emas -> ", original)
	}
}

func main() {
	palindromSon(12321)
}