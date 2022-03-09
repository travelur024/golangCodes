package main

import (
	"fmt"
	"time"
)

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}

func main() {
	timer1 := time.Now()
	// //Creating a channel
	// for i := 0; i < 20; i++ {
	// 	fmt.Println("Done")
	// }
	// c := make(chan int)
	// go doSomething(c)
	// <-c

	//Apuntadores
	original := 25
	fmt.Println("Valor orignal", original)
	//apuntador
	copia1 := &original
	fmt.Println("Apuntador copia1", *copia1, original)
	*copia1 = 334
	var copia2 *int = copia1
	fmt.Println("Reasignacion a 334 copia1:", *copia1, *copia2, original)
	*copia2 = 69
	fmt.Println("Reasignacion a 69 copia2:", *copia1, *copia2, original)
	original = 45
	fmt.Println("Reasignacion a 45 origianal:", *copia1, *copia2, original)
	elapsed := time.Since(timer1)
	fmt.Println(elapsed)

}
