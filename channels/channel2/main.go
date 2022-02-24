package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimir(ch chan<- string, texto string) {
	for {
		// Pasando información al channel.
		ch <- texto
		// Tiempo de espera entre cada iteración.
		// rand.Int nos permite generar un entero de manera aleatoria
		// time.Sleep requiere que la variable sea de tipo time.Duration
		time.Sleep(time.Duration(rand.Int()) * time.Nanosecond)
	}

}

func main() {
	// Creando un channel de tipo string.
	ch := make(chan string)

	// Pasando el channel a dos goroutines.
	go imprimir(ch, "primera")
	go imprimir(ch, "segunda")
	go imprimir(ch, "tercera")
	go imprimir(ch, "cuarta")
	go imprimir(ch, "quinta")
	//close(ch)

	// Ciclo infinito que imprime la información que viene en el channel.
	for {

		valor, ok := <-ch
		if ok != true {
			break
		}
		fmt.Println(valor)

		// for message := range ch {
		// 	fmt.Println(message)
		// }

	}
}
