package main

import (
	"fmt"

	"github.com/travelur024/003BasicCourse/reto/rt"
)

func main() {
	myPc := rt.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}
