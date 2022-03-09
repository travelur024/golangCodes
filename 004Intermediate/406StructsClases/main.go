package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{} //Creando instancia
	fmt.Println(e)
	e.id = 86
	e.name = "Jonathan"
	fmt.Println(e)
}
