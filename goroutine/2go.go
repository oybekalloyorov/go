package main

import (
	"fmt"
	"sync"
	"time"
)

type SonMalumot struct {
	AslSon   int
	Kvadrati int
	Juftmi   bool
}

func sonniQaytaIshla(son int, wg *sync.WaitGroup, natijaCh chan<- SonMalumot) {
	defer wg.Done()
	
}
func ishla(son int,wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Printf(" Son: %d, Kvadrati: %d, Juftmi: %t\n",
	// 	son, son*son,
	// 	son%2 == 0,
	// )
}

func main() {
	now := time.Now()
	sonlar := []int{1, 2, 3, 4, 5,6,7,8,9,10}
	var wg sync.WaitGroup
	// results := make(chan SonMalumot, len(sonlar))
	for _, son := range sonlar {
		wg.Add(1)
		 ishla(son, &wg)
	}

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		 ishla(i, &wg)
	}
	wg.Wait()
	fmt.Println("Umumiy vaqt:", time.Since(now))


}