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
	fmt.Println(time.Now())
	buyurtmalar := []Buyurtma{
		{1,  []Taom{Salat, Shorva, AsosiyTaom}},
	}

	for _, buyurtma := range buyurtmalar {
		for _, taom := range buyurtma.Taomlar {
			tayyorlash(taom)
		}
	}
	// fmt.Println("Barcha buyurtmalar tayyor!")

}
