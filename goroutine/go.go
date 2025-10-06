package main

import (
	"fmt"
	"sync"
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

func tayyorlash(wg *sync.WaitGroup, taom Taom) {
	defer wg.Done()
	fmt.Printf("%s tayyorlanmoqda...\n", taom.Nomi)
	time.Sleep(taom.TayyorlashVaqti)
	fmt.Printf("%s tayyor!\n", taom.Nomi)
}

func main() {
	var wg sync.WaitGroup

	now := time.Now()
	buyurtmalar := []Buyurtma{
		{1,  []Taom{Salat, Shorva, AsosiyTaom}},
		{2,  []Taom{Salat, Shorva, AsosiyTaom}},
		{3,  []Taom{Salat, Shorva, AsosiyTaom}},
		{4,  []Taom{Salat, Shorva, AsosiyTaom}},
	}
	total := 0
	for _, b := range buyurtmalar{
		total += len(b.Taomlar)
	}
	wg.Add(total)

	for _, buyurtma := range buyurtmalar {
		for _, taom := range buyurtma.Taomlar {
			go tayyorlash(&wg, taom)
		}
	}
	// time.Sleep(6 * time.Second) // Barcha taomlar tayyor bo'lishi uchun kutamiz
	wg.Wait()
	fmt.Println("Umumiy vaqt:", time.Since(now))

}
