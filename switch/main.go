package main

import "fmt"

func main() {
	fmt.Println("Write a number:")
	var num1 int
	fmt.Scanln(&num1)
	modulo := num1 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")

	}

	whaType(3.14)
	whaType(4134)
	whaType("Go")

}

func whaType(value interface{}) {

	switch v := value.(type) {
	case float64:
		fmt.Printf("%g is float. \n", v)
	case int:
		fmt.Printf("%d is integer. \n", v)
	default:
		fmt.Printf("%v is %T. \n", v, v)
	}
}
