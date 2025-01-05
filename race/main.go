package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.GOMAXPROCS(0))

	contador := 0
	totalGoRoutines := 100
	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	var mu sync.Mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			mu.Lock()
			count := contador
			runtime.Gosched()
			count++
			wg.Done()
			contador = count
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Contador", contador)
}
