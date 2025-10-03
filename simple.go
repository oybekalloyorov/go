// package main


// import (
// "bytes"
// "fmt"
// "io"
// )


// func ReadAll(r io.Reader) string {
// buf := make([]byte, 1024)
// n, _ := r.Read(buf)
// return string(buf[:n])
// }


// func main() {
// r := bytes.NewBufferString("Salom, dunyo")
// fmt.Println(ReadAll(r))
// }

package main


import (
"fmt"
"sort"
)


type Person struct { Name string; Age int }


type ByAge []Person


func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }


func main() {
people := []Person{{"Ali", 30}, {"Vali", 25}, {"Olim", 40}}
sort.Sort(ByAge(people))
fmt.Println(people)
}
