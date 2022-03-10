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

func (ft FullTimeEmployee) GetMessage() string {
	return "Full time employee" + ft.name
}

type TemporaryEmployee struct {
	Persona
	Employee
	taxRate int
}

func (te TemporaryEmployee) GetMessage() string {
	return "Temporary employee"
}

type Printlnfo interface {
	GetMessage() string
}

func GetMessage(p Printlnfo) {
	fmt.Println(p.GetMessage())
}

func main() {
	//En go existe la composicion sobre herencia
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Jonathan"
	ftEmployee.edad = 24
	ftEmployee.id = 5
	fmt.Println(ftEmployee)
	tEmployee := TemporaryEmployee{}
	GetMessage(tEmployee)
	GetMessage(ftEmployee)

	//GetMessage(ftEmployee)		//Wrong

}
