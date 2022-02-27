package main

import "fmt"

//Animales que indican como se estan moviendo - Ejemplo interfaces
//Structs basicos de animales
type perro struct {
}

type pez struct {
}

type pajaro struct {
}

//Metodos de animales- Con la accion que necesitamos
func (perro) caminar() string {
	return "Soy un perro y camino"
}
func (pez) nada() string {
	return "Soy un pez y nado"
}
func (pajaro) vuela() string {
	return "Soy un pajaro y vuelo"
}

//Metodo que imprima que esta ocurriendo como se mueve el animal
func moverPerro(p perro) {
	fmt.Println(p.caminar())
}
func moverPez(p pez) {
	fmt.Println(p.nada())
}
func moverPajaro(p pajaro) {
	fmt.Println(p.vuela())
}
func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	moverPez(pe)
	pa := pajaro{}
	moverPajaro(pa)
}
