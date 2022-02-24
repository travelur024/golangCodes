package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	//Channel - Un conducto que maneja un tipo de dato
	c := make(chan string, 1)
	fmt.Println("Hello")

	go say("bye", c)

	fmt.Println(<-c)
}
