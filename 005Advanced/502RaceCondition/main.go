package main

import (
	"errors"
	"fmt"
	"sync"
)

var balance int

func Deposit(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	balance = value + balance
	fmt.Println("Deposito:", balance)
}

func Withdraw(value int, wg *sync.WaitGroup) error {
	defer wg.Done()
	if value > balance {
		return errors.New("You can't do that man!")
	}
	balance = balance - value
	fmt.Println("Retiro:", balance)
	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	//WaitGroup

	balance = 500
	go Withdraw(300, &wg)
	go Deposit(400, &wg)
	wg.Wait()
	fmt.Println(balance)

}
