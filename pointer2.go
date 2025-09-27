package main

import "fmt"

type stack struct {
    index int
    data  [5]int
}

// func (s stack) push(k int) {
//     s.data[s.index] = k
//     s.index++
// }
func (s stack) push(k int) stack {
    s.data[s.index] = k
    s.index++
    return s
}

func (s stack) pop() int {
    s.index--
    return s.data[s.index]
}

func main() {
    // s := new(stack) // *stack
    // s.push(23)
    // s.push(14)
    // s.push(19)
    // fmt.Printf("stack: %v\n", *s) // NATIJA: {0 [0 0 0 0 0]}

	var s stack
    s = s.push(23);
    s = s.push(14);
    s = s.push(19);
    fmt.Printf("stack: %v\n", s) // {3 [23 14 19 0 0]}


}
