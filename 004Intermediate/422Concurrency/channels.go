package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	fmt.Println("Start")
	go func() {
		c <- 1
		c <- 2
		c <- 3

	}()
	go func() {
		fmt.Println(<-c)
		time.Sleep(1 * time.Second)
		fmt.Println(<-c)
		time.Sleep(1 * time.Second)
		fmt.Println(<-c)
		fmt.Println("End")
	}()
	time.Sleep(5 * time.Second)

}
