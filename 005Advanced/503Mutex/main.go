package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

//Llamando a lock en la funcion que bloquea la variable a modificar
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup
	//Lock
	var lock sync.Mutex
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(balance)
}
