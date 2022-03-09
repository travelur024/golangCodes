package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (this *Employee) SetId(id int) {
	this.id = id
}

func (this *Employee) SetName(name string) {
	this.name = name
}

func (this *Employee) GetId() int {
	return this.id
}

func (this *Employee) GetName() string {
	return this.name
}

func main() {
	e := Employee{} //Creando instancia
	fmt.Println(e)
	e.id = 69
	e.name = "Jonathan"

	fmt.Println(e)
	e.SetId(19)
	e.SetName("Jose Jose")
	fmt.Println(e)
	fmt.Println("Usando metodos get")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	//TODO deber ser un objeto (filosofia no siempres aplicable)
	//Ejemplo: En una libreria con utilidades que no pertenecen
	//a un dominio especifico.
}
