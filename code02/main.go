package main

import (
	"fmt"
)

func areaRectangulo(base int, altura int) int {
	return base * altura
}

func main() {
	//Declaraicon de constantes
	//Forma1
	const pi float64 = 3.1416
	//Forma2
	const pi2 = 3.15

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 24
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const base01 int = 10
	const altura01 = 3
	fmt.Println(base01, altura01)
	fmt.Println("Area:", (areaRectangulo(base01, altura01)))

	// modification
}
