package main

import (
	"fmt"

)

func main(){
	// var n int
	// fmt.Print("Son kiriting: ")
	// fmt.Scan(&n)
	// if n%2 == 0 {         //Juft yoki Toq
	// 	fmt.Println(n, "Juft")
	// }else{
	// 	fmt.Println(n, "Toq")
	// }

	//Yigindi hisoblash
	// sum := 0
	// for i := 1; i <= n; i++ {
	// 	sum += i
	// }
	// fmt.Println("1..", n, "yigindisi =", sum)

	// x, y := swap("hello", "world")
	// fmt.Println(x, y)

	// Factarial
	// fmt.Println("6! =", fact(6))

	// Oddiy calculyator 
	var a, b int
	var op string

	fmt.Print("a: "); fmt.Scan(&a)
	fmt.Print("b: "); fmt.Scan(&b)
	fmt.Print("amal (+ - * /): "); fmt.Scan(&op)

	switch op{
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b != 0 {
			fmt.Println(a / b)
		}else {
			fmt.Println("0 ga bo'lib bo'lmaydi")
		}
	default:
		fmt.Println("Noma'lum amal")
	}
}
	//Swap funksiya
	func swap(a, b string)(string, string)  {
		return b, a
	}

	// Factarial
	func fact(n int) int{
		if  n <= 1 {
			return 1
		}
		return n * fact(n-1)
	}