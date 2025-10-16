package main

import "fmt"

func main(){
	slice := []int{1, 2, 3}

	slice = append(slice, 4)
	slice = append(slice, 5)

	fmt.Println(slice)

	//28
	// a := []int{1,2,3}
	// b := a          // b a bilan bir xil underlying array'ga ishora qiladi
	// b[0] = 99
	// fmt.Println(a)  // natija: [99 2 3]
	// fmt.Println(b)  // natija: [99 2 3]

	slice2 := []int{1,2,3}
	copySlice := make([]int, len(slice2)) // yoki make([]int, len(src), len(src))
	copy(copySlice, slice2)
	copySlice[0] = 99

	fmt.Println("src:", slice2) // [1 2 3]
	fmt.Println("dst:", copySlice) // [99 2 3]

}