package main

import (
	"fmt"
	"time"
)

type Taom struct {
	Nomi  string
	TayyorlashVaqti time.Duration
}

var (
	Salat      = Taom{"Salat", 2 * time.Second}
	Shorva     = Taom{"Shorva", 3 * time.Second}
	AsosiyTaom = Taom{"Asosiy Taom", 5 * time.Second}
)

type Buyurtma struct {
	ID      int
	Taomlar []Taom
}

func tayyorlash(taom Taom) {
	fmt.Printf("%s tayyorlanmoqda...\n", taom.Nomi)
	time.Sleep(taom.TayyorlashVaqti)
	fmt.Printf("%s tayyor!\n", taom.Nomi)
}

func main() {
	now := time.Now()
	buyurtmalar := []Buyurtma{
		{1,  []Taom{Salat, Shorva, AsosiyTaom}},
		{2,  []Taom{Salat, Shorva, AsosiyTaom}},
		{3,  []Taom{Salat, Shorva, AsosiyTaom}},
		{4,  []Taom{Salat, Shorva, AsosiyTaom}},
	}

	for _, buyurtma := range buyurtmalar {
		for _, taom := range buyurtma.Taomlar {
			go tayyorlash(taom)
		}
	}
	time.Sleep(6 * time.Second) // Barcha taomlar tayyor bo'lishi uchun kutamiz
	fmt.Println("Umumiy vaqt:", time.Since(now))

}
