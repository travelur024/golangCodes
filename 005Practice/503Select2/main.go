package main

import (
	"fmt"
)

//Funcion que resive datos indefinidamente...
func ReciverDataFromChannel(a int, c chan<- int, out <-chan int) {
	count := a
	for {

		select {
		case c <- count:
			fmt.Println("Recibi: ", count)
		case <-out:
			fmt.Println("quit")
			return

		}
		count++
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//Antes de escribir al canal es necesario leerlo
	// Leemos 2 veces
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Out of channel: ", <-c)
		}
		quit <- 0
	}()
	ReciverDataFromChannel(0, c, quit)
}
