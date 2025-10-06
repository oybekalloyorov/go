package main

import (
	"fmt"
	"sync"
)

type Hisob struct {
	son int
	mu      *sync.Mutex
}

func (h *Hisob)oshirish(wg *sync.WaitGroup) {
	defer wg.Done()
	h.mu.Lock()
	h.son++
	h.mu.Unlock()
}

func main() {
	
	var wg sync.WaitGroup
	h := Hisob {
		son: 0,
		mu:      &sync.Mutex{},
	}

	goroutineSoni := 1000

	for i := 0; i < goroutineSoni; i++ {
		wg.Add(1)
		go h.oshirish(&wg)
	}
	wg.Wait()
	fmt.Println("Final hisoblagich qiymati:", h.son)
}