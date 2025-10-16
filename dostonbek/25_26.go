package main

import "fmt"

func juftSonlarSoni(a [5]int) int{
	count := 0
	for i := 0; i < len(a); i++ {
		if (a[i] % 2 == 0) {
			count++
		}
	}
	return count
}

func massivniTeskariChiqarish(a [5]int){
	for i := len(a)-1; i >= 0; i-- {
		fmt.Print(a[i] , " ")
	}
}

func main(){
	arr := [5]int{2,3,4,5,6}

	fmt.Println("massivdagi juft sonlar soni = ", juftSonlarSoni(arr))

	massivniTeskariChiqarish(arr)
}