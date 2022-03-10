package main

import "fmt"

type Persona struct {
	name string
	edad int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Persona
	Employee
}

func GetMessage(p Persona) {
	fmt.Printf("%s with age %d\n", p.name, p.edad)
}

func main() {
	//En go existe la composicion sobre herencia
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Jonathan"
	ftEmployee.edad = 24
	ftEmployee.id = 5
	fmt.Println(ftEmployee)
	//GetMessage(ftEmployee)		//Wrong

}
