package main

import (
	"fmt"
	"time"
)

func main() {

	// make the request chan chan that both go-routines will be given
	requestChan := make(chan chan string)

	// start the goroutines
	go goroutineC(requestChan)
	go goroutineD(requestChan)

	// sleep for a second to let the goroutines complete
	time.Sleep(time.Second)

}

func goroutineC(requestChan chan chan string) {

	// make a new response chan
	responseChan := make(chan string)
	//responseChan2 := make(chan string)

	// send the responseChan to goRoutineD
	requestChan <- responseChan
	// requestChan <- responseChan2

	// read the response
	response := <-responseChan
	// response2 := <-responseChan2
	// amiko := <-responseChan2

	fmt.Printf("Response: %v\n", response)
	// fmt.Printf("Response2: %v and How writting: %v\n", response2, amiko)
}

func goroutineD(requestChan chan chan string) {

	// read the responseChan from the requestChan
	responseChan := <-requestChan
	// responseChan2 := <-requestChan

	// send a value down the responseChan
	responseChan <- "wassup!"
	// responseChan2 <- "002! prro"
	// responseChan2 <- "amikooo?"

}
