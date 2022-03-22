package main

import "fmt"

func fibonacci(c chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("Recibing: ", x, y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//Nesario solicitar ya que manejamos unbuffered
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	//Iniciado la solicitud, es poosible mandar datos...
	fibonacci(c, quit)
}
