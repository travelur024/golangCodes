package main

import "fmt"

func main() {
	slice := []string{"hola", "khe", "ace"}
	for _, elem := range slice {
		fmt.Println(elem)
	}
}
