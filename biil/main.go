package main

import "fmt"

func main() {
	mybill := newBill("marios bill")

	fmt.Println(mybill.format())
}