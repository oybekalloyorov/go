package main

import "sync"

func main() {
	var num int
	var slice []int

	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i:= range 1_000_000 {
			mx.Lock()
			slice = append(slice, i)
			num++
			mx.Unlock()
		}
		println("gr 1 is finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i:= range 1_000_000 {
			mx.Lock()
			slice = append(slice, i)
			num++
			mx.Unlock()
		}
		println("gr 2 is finished")
	}()

	wg.Wait()
	println("main is finished", num, "\nslice -> ", len(slice))
}