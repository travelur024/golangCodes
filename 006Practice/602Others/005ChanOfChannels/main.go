package main

import "fmt"

func main() {
	//Canal de canales
	values := []int{1, 9, 3, 4, 5, 7, 8}
	c := make(chan int, len(values))
	mchan := make(chan chan int, 4)

	for _, elem := range values {
		go Receving(c, elem)
	}
	//Masterchannel
	for i := 0; i < 3; i++ {
		go RecevingMC(mchan, c)
	}

	// Mando a imprimir la data que se mando...
	// ////////////////////////////////////////////////////////////////////
	for i := 0; i < 2; i++ {
		data := <-mchan
		for i := 0; i < len(values); i++ {
			fmt.Println("Values", <-data)
		}
		fmt.Println("Data MC: ", data)
	}

	//data<- mchan
	// for i := 0; i < len(values); i++ {
	// 	fmt.Println("Data from MChannel 01:", <-data)
	// }

	////////////////////////////////////////////////////////////////////
	// fmt.Println(<-mchan)
	// fmt.Println(<-mchan)
	//////////////}

	// for i := 0; i < 2; i++ {
	// 	fmt.Println(<-mchan)
	// }

	// for i := 0; i < len(values); i++ {
	// 	fmt.Println("Lastchange:", <-c)
	// }

	// for i := 0; i < len(values); i++ {
	// 	fmt.Println(<-c)
	// }

}

func Receving(c chan<- int, a int) {
	c <- a
}

func RecevingMC(mc chan chan int, c chan int) {
	mc <- c
}
