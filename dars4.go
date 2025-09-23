package main

import (
    "fmt"
    "math"
)

// Shape interface: hamma shakl uchun umumiy metodlar
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle - value receiver metodlar
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle - pointer receiver metodlari
type Circle struct {
    Radius float64
}

func (c *Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
func (c *Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Scale modifies rectangle (pointer receiver)
func (r *Rectangle) Scale(f float64) {
    r.Width *= f
    r.Height *= f
}

func main() {
    // Diqqat: Circle metodlari pointer receiver bo'lgani uchun &Circle yaratish tavsiya etiladi
    shapes := []Shape{
        Rectangle{Width: 3, Height: 4}, // value type is OK for Rectangle
        &Circle{Radius: 2},             // pointer required for Circle
    }

    for _, s := range shapes {
        fmt.Printf("Area: %.2f, Perim: %.2f\n", s.Area(), s.Perimeter())
    }

    // Scale qilish misoli
    r := Rectangle{Width: 2, Height: 5}
    r.Scale(2) // value r ni pointerga avtomatik o'zgartirib chaqirish mumkin emas — bu yerda compiler r.Scale(2) kompilyatsiya qiladi; (r is addressable)
    fmt.Println("Scaled rectangle:", r)
}
