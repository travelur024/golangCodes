package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)

}

func main() {

	//waitGroup
	var wg sync.WaitGroup

	fmt.Println("Hellow")
	wg.Add(5)
	go say("World1", &wg)
	go say("World2", &wg)
	go say("World3", &wg)
	go say("World4", &wg)
	go say("World5", &wg)
	wg.Wait()

	//Funcion anonima
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	//Nesesario pero no es buena practica
	time.Sleep(time.Second * 1)

}
