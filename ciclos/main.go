package main

import "fmt"

func main() {

	//For condicional (easy)
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n.") //separacion

	//For while

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for i, elem := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, elem)
	}

	//	For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 3 {
			break
		}
	}

}
