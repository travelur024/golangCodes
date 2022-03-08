package main

import "fmt"

func main() {
	//Operadores logicos..
	//Forma1
	const pi float64 = 3.1416
	const c1 int32 = 23
	const c2 int16 = 59
	const c3 int8 = 34
	//Forma2
	const pi2 = 3.15
	const real = 69.234
	const real2 = 34.123
	const real3 = 67.432
	//Variable
	var second string
	fmt.Scanf("%s", &second)
	fmt.Println(second)
	//Operador logico
	if pi != 3.1416 && c1 == 23 {
		fmt.Println("pi is INCORRECT!", pi)
	} else {
		fmt.Println("pi is CORRECT!", pi)
	}

}
