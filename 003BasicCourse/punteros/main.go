package main

import "fmt"

func main() {
	a := 50 //Variable a
	b := &a //Direccion de memoria de a

	fmt.Println(b)  //Direccion de memoria
	fmt.Println(*b) //Valor de la direccion de memoria

	//Si se modifica el valor que apunta a la direccion
	//de memoria todas la variables que apuntan a esa direccion
	//cambian
	*b = 69

	fmt.Println(a)  //Direccion de memoria
	fmt.Println(*b) //Valor de la direccion de memoria

}
