package main

import "fmt"

func main() {
	//Creamos la fabrica
	FactoryAdidas, _ := getSportsFactory("adidas")
	ShirtAdidas001 := FactoryAdidas.makeShirt()
	ShoeAdidas001 := FactoryAdidas.makeShoe()
	ShirtAdidas001.setLogo("MDXZS")
	ShoeAdidas001.setBrand("MDXZ002")
	ShoeAdidas001.setSize(19)

	printShoeDetails(ShoeAdidas001)
	printShirtDetails(ShirtAdidas001)

}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getBrand())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
