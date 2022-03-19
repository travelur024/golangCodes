package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

//Llamando a lock en la funcion que bloquea la variable a modificar
func Deposit(amount int, wg *sync.WaitGroup, mux *sync.RWMutex) {
	defer wg.Done()
	mux.Lock()
	b := balance
	balance = b + amount
	mux.Unlock()
}

//1 Deposit() -> Escribiendo (Race Condition)
//N Balance() -> Leer

func Balance(mux *sync.RWMutex) int {
	mux.RLock()
	b := balance
	mux.RUnlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	//Lock
	var lock sync.RWMutex //lock de escritura y lectura
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
