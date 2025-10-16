package main

import "fmt"

func massivdagiEngKattaSon(arr [4]int) int{
	kattaSon := arr[0]

	for i := 0; i < len(arr); i++ {
		if(kattaSon < arr[i]){
			kattaSon = arr[i]
		}
	}
	return kattaSon
}

func main() {
	massiv := [4]int{1,2,3,4}

	sum := 0
	for i := 0; i < len(massiv); i++ {
		sum += massiv[i]
	}

	fmt.Println(sum)

	fmt.Println(massivdagiEngKattaSon(massiv))

}