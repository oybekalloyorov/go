package main

func main() {
	// sum(1,2,3,4,5,6,7,8,9,10)
	summa, _ := sum(10,20,30,40,50,1)
	println(summa)
}

func sum(a ... int) (int, string){
	var summa int
	for _, i := range a {
		summa += i
	}
	return summa, "Yakunlandi"
}