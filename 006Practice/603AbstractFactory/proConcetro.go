package main

//Productos concretos

type adidasShoe struct {
	shoe
}
type nikeShoe struct {
	shoe
}

type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}

// gitSize implements iShirt
func (*adidasShirt) gitSize() int {
	panic("unimplemented")
}

// gitSize implements iShirt
func (*nikeShirt) gitSize() int {
	panic("unimplemented")
}
