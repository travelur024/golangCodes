package main

import (
	"fmt"

	pk "github.com/travelur024/accesos/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Nissan"
	myCar.Year = 2020

	fmt.Println(myCar)

}
