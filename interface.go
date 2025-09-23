package main

import (
	"fmt"
	"math"
)

type Animal interface{
	Speak() string
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func PrintAnything(x interface{}){
	fmt.Println(x)
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string{
	return "woof!"
}

func (c Cat) Speak() string{
	return "Meow!"
}

func main(){
	var a Animal

	a = Dog{}

	fmt.Println(a.Speak())

	a = Cat{}

	fmt.Println(a.Speak())

	c := Circle{Radius: 5}

	fmt.Println(c.Perimeter())
	fmt.Println(c.Area())

	
	PrintAnything(42)
	PrintAnything("Salom")
    PrintAnything(true)

}