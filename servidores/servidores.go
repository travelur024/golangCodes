package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//Timer
	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	//Tiempo final
	tiempoPaso := time.Since(inicio)
	fmt.Println(tiempoPaso)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta disponible =(")
	} else {
		fmt.Println(servidor, "Esta funcionando normalmente =)")
	}
}
