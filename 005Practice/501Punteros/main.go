package main

import "fmt"

type V struct {
	i int
}

func incrementa(v *V) {
	v.i++
	fmt.Println(v.i)
}

func inc002(a *int) {
	*a++
}

func main() {
	a := V{i: 25}
	incrementa(&a)
	fmt.Println(a)
	//Creando puntero
	//Puntero: variable que almacena la direccion
	//en memoria de un valor

	//Creando puntero
	b := &a.i
	fmt.Printf("Valor de b: %v\n", b)
	fmt.Printf("Direc de b: %v\n", &b)
	//Acceder al valor que apunta el puntero
	//Dereferencia el puntero *nombrePuntero
	fmt.Printf("Valor al que apunta b: %v\n", *b)

	//Modificar el valor del punteto
	*b = 27
	fmt.Printf("Valor al que apunta b: %v  Valor original a: %v \n", *b, a.i)
	//Using other function with the pointer that contents a dictionary of memori &a...
	inc002(b)
	fmt.Printf("Valor al que apunta b: %v  Valor original a: %v \n", *b, a.i)

}
