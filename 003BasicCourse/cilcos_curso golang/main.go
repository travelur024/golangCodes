package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//Timer
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	//leyendo el canal
	// fmt.Println(<-canal) //Solo espera una vez al canal y terman el main
	//Contador
	i := 0

	for {
		if i > 3 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	//Tiempo final
	tiempoPaso := time.Since(inicio)
	fmt.Println(tiempoPaso)
}

func revisarServidor(servidor string, chanel chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		chanel <- servidor + "no se encuentra disponible =("
	} else {
		chanel <- servidor + "esta funcionando normalmente =)"
	}
}
