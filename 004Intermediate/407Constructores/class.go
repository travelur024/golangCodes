package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//1
	e := Employee{}
	fmt.Println(e)
	//2
	e2 := Employee{
		id:       234,
		name:     "Jonathan",
		vacation: true,
	}
	fmt.Println(e2)
	//3
	e3 := new(Employee)
	e3.id = 69
	e3.name = "Raurul Chessel"
	e3.vacation = true
	fmt.Println(*e3)
	//4
	e4 := NewEmployee(24, "Name05", false)
	fmt.Println(*e4)
	e4 = NewEmployee(69, "Ramirito", true)
	fmt.Println(*e4)
}
