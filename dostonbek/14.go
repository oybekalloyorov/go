package main

import "fmt"

func teskariRaqamYasash(n int) {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	fmt.Println(reversed)
}

func main() {
	teskariRaqamYasash(12345)
}