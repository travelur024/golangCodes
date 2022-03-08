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

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	//leyendo el canal
	// fmt.Println(<-canal) //Solo espera una vez al canal y terman el main

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	//Tiempo final
	tiempoPaso := time.Since(inicio)
	fmt.Println(tiempoPaso)
}

func revisarServidor(servidor string, chanel chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta disponible =(")
		chanel <- servidor + "no se encuentra disponible =("
	} else {
		fmt.Println(servidor, "Esta funcionando normalmente =)")
		chanel <- servidor + "esta funcionando normalmente =)"
	}
}
