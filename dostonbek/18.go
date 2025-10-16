// package main

// import "fmt"

// func main() {
//     x := 10
//     p := &x
//     fmt.Println("x:", x, "p:", p, "*p:", *p)
//     *p = 20
//     fmt.Println("x after *p=20:", x)
// }

// package main

// import "fmt"

// type Person struct {
//     Name string
//     Age  int
// }

// func (p *Person) HaveBirthday() {
//     p.Age++
// }

// func main() {
//     person := &Person{Name: "Ali", Age: 30}
//     person.HaveBirthday()
//     fmt.Println(person.Age) // 31
// }

package main

import "fmt"

func main() {
    p := 10

	p2 := &p
	fmt.Println("p:", p, "p2:", p2, "*p2:", *p2)
	*p2 = 20
	fmt.Println("p after *p2=20:", p)
}

