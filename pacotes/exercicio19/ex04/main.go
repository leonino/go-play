package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const totalGoRoutines = 100

var mu sync.Mutex

func main() {
	criarGoRoutines()
	wg.Wait()
	fmt.Println("Total de GoRoutines: ", totalGoRoutines, "Contador: ", contador)
}

func criarGoRoutines() {
	wg.Add(totalGoRoutines)
	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
