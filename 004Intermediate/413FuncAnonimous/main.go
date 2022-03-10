package main

import (
	"fmt"
	"time"
)

//func double

func main() {
	//Funcion anonima
	// x := 5
	// y := func() int {
	// 	return x * 2
	// }()
	// fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("Iniciando la funcion")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	go func() {
		fmt.Println("Iniciando la funcion")
		time.Sleep(3 * time.Second)
		fmt.Println("End")
		c <- 1
	}()

	<-c
	<-c
}
