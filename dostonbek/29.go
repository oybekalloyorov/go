package main

import "fmt"

func slicedanQidir(slice []int, qidir int)(bool){
	topildi := false 
	for _, v := range slice{
		if (v == qidir) {
			topildi = true
			break
		}
	}

	return topildi
}

func main(){
	slice := []int{1, 4, 7}
	qidir := 4
	res := slicedanQidir(slice, qidir)

	if (res) {
		fmt.Println("Topildi")
	}else{
		fmt.Println("Topilmadi")
	}
		
}
