package main

import "fmt"

//Crear la interface superclase
type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

//Factory clase base
type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}
func (c *Computer) setName(name string) {
	c.name = name
}
func (c *Computer) getName() string {
	return c.name
}
func (c *Computer) getStock() int {
	return c.stock
}

//Productos - Forma en que funcion la compocision sobre la herencia
type Laptop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

//Type descktop

type Desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &Desktop{
		Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

//Creacional - Factory

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}
	if computerType == "desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("Invalid computer type!")
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")
	printNameAndStock(laptop)
	printNameAndStock(desktop)
}

//Clase base computer// clases derivadas latop desktop que cambian los comportamientos 