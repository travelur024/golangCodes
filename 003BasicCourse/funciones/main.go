package main

import "fmt"

func mensajeConcatenado(msj string) string {
	fmt.Println(msj)
	return "Jokei impreso bro"
}

func main() {
	var done string
	done = mensajeConcatenado("Helada la noche pero no tanto como el invierno")
	fmt.Println(done)
}
