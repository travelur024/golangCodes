package main

import "fmt"

func main() {
	//Map estructura Llave - Valor - String a numero
	m1 := make(map[string]int)

	m1["a"] = 8
	fmt.Println(m1)
	fmt.Println(m1["a"])
}
