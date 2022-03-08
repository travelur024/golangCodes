package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	//Primera forma de instanciar un Struct
	myCar := car{brand: "Nissan", year: 2020}
	fmt.Println(myCar)

	//Other form
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
