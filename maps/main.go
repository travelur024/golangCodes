package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	m1 := make(map[string]bool)
	m1["Jose"] = true
	m1["Pepito"] = false

	fmt.Println(m1)

	//Recorreindo un Range
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor in maps
	value, ok := m["Jose"]
	fmt.Println(value, ok)

}
