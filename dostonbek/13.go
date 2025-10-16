package main

import "fmt"

func sonRaqamYigindisi(n int) {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	fmt.Println(sum)
}

func main() {
	sonRaqamYigindisi(12345)
}