package main

import "fmt"

func slicedanOchirish(slice []int, del int)([]int){
	for i, v := range slice{
		if (v == del){
			slice = append(slice[:i], slice[i+1:]... )
			return slice
		}
	}
	return slice
}

func main(){
	slice := []int{10, 20, 30, 40}
	del := 300

	newSlice := slicedanOchirish(slice, del)
	fmt.Println(newSlice)
}